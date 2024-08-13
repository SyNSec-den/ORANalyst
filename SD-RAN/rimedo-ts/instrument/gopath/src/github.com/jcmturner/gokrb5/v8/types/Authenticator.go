//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:1
// Package types provides Kerberos 5 data types.
package types

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:2
)

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/asn1tools"
	"github.com/jcmturner/gokrb5/v8/iana"
	"github.com/jcmturner/gokrb5/v8/iana/asnAppTag"
)

// Authenticator - A record containing information that can be shown to have been recently generated using the session
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:17
// key known only by the client and server.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:17
// https://tools.ietf.org/html/rfc4120#section-5.5.1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:20
type Authenticator struct {
	AVNO			int			`asn1:"explicit,tag:0"`
	CRealm			string			`asn1:"generalstring,explicit,tag:1"`
	CName			PrincipalName		`asn1:"explicit,tag:2"`
	Cksum			Checksum		`asn1:"explicit,optional,tag:3"`
	Cusec			int			`asn1:"explicit,tag:4"`
	CTime			time.Time		`asn1:"generalized,explicit,tag:5"`
	SubKey			EncryptionKey		`asn1:"explicit,optional,tag:6"`
	SeqNumber		int64			`asn1:"explicit,optional,tag:7"`
	AuthorizationData	AuthorizationData	`asn1:"explicit,optional,tag:8"`
}

// NewAuthenticator creates a new Authenticator.
func NewAuthenticator(realm string, cname PrincipalName) (Authenticator, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:33
	_go_fuzz_dep_.CoverTab[85943]++
													seq, err := rand.Int(rand.Reader, big.NewInt(math.MaxUint32))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:35
		_go_fuzz_dep_.CoverTab[85945]++
														return Authenticator{}, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:36
		// _ = "end of CoverTab[85945]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:37
		_go_fuzz_dep_.CoverTab[85946]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:37
		// _ = "end of CoverTab[85946]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:37
	// _ = "end of CoverTab[85943]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:37
	_go_fuzz_dep_.CoverTab[85944]++
													t := time.Now().UTC()
													return Authenticator{
		AVNO:		iana.PVNO,
		CRealm:		realm,
		CName:		cname,
		Cksum:		Checksum{},
		Cusec:		int((t.UnixNano() / int64(time.Microsecond)) - (t.Unix() * 1e6)),
		CTime:		t,
		SeqNumber:	seq.Int64(),
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:47
	// _ = "end of CoverTab[85944]"
}

// GenerateSeqNumberAndSubKey sets the Authenticator's sequence number and subkey.
func (a *Authenticator) GenerateSeqNumberAndSubKey(keyType int32, keySize int) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:51
	_go_fuzz_dep_.CoverTab[85947]++
													seq, err := rand.Int(rand.Reader, big.NewInt(math.MaxUint32))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:53
		_go_fuzz_dep_.CoverTab[85949]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:54
		// _ = "end of CoverTab[85949]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:55
		_go_fuzz_dep_.CoverTab[85950]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:55
		// _ = "end of CoverTab[85950]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:55
	// _ = "end of CoverTab[85947]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:55
	_go_fuzz_dep_.CoverTab[85948]++
													a.SeqNumber = seq.Int64()

													sk := make([]byte, keySize, keySize)
													rand.Read(sk)
													a.SubKey = EncryptionKey{
		KeyType:	keyType,
		KeyValue:	sk,
	}
													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:64
	// _ = "end of CoverTab[85948]"
}

// Unmarshal bytes into the Authenticator.
func (a *Authenticator) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:68
	_go_fuzz_dep_.CoverTab[85951]++
													_, err := asn1.UnmarshalWithParams(b, a, fmt.Sprintf("application,explicit,tag:%v", asnAppTag.Authenticator))
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:70
	// _ = "end of CoverTab[85951]"
}

// Marshal the Authenticator.
func (a *Authenticator) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:74
	_go_fuzz_dep_.CoverTab[85952]++
													b, err := asn1.Marshal(*a)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:76
		_go_fuzz_dep_.CoverTab[85954]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:77
		// _ = "end of CoverTab[85954]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:78
		_go_fuzz_dep_.CoverTab[85955]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:78
		// _ = "end of CoverTab[85955]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:78
	// _ = "end of CoverTab[85952]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:78
	_go_fuzz_dep_.CoverTab[85953]++
													b = asn1tools.AddASNAppTag(b, asnAppTag.Authenticator)
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:80
	// _ = "end of CoverTab[85953]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:81
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Authenticator.go:81
var _ = _go_fuzz_dep_.CoverTab
