//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:1
package types

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:1
)

import (
	"crypto/rand"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/crypto/etype"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:13
// EncryptedData implements RFC 4120 type: https://tools.ietf.org/html/rfc4120#section-5.2.9
type EncryptedData struct {
	EType	int32	`asn1:"explicit,tag:0"`
	KVNO	int	`asn1:"explicit,optional,tag:1"`
	Cipher	[]byte	`asn1:"explicit,tag:2"`
}

// EncryptionKey implements RFC 4120 type: https://tools.ietf.org/html/rfc4120#section-5.2.9
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:20
// AKA KeyBlock
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:22
type EncryptionKey struct {
	KeyType		int32	`asn1:"explicit,tag:0"`
	KeyValue	[]byte	`asn1:"explicit,tag:1" json:"-"`
}

// Checksum implements RFC 4120 type: https://tools.ietf.org/html/rfc4120#section-5.2.9
type Checksum struct {
	CksumType	int32	`asn1:"explicit,tag:0"`
	Checksum	[]byte	`asn1:"explicit,tag:1"`
}

// Unmarshal bytes into the EncryptedData.
func (a *EncryptedData) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:34
	_go_fuzz_dep_.CoverTab[85959]++
													_, err := asn1.Unmarshal(b, a)
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:36
	// _ = "end of CoverTab[85959]"
}

// Marshal the EncryptedData.
func (a *EncryptedData) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:40
	_go_fuzz_dep_.CoverTab[85960]++
													edb, err := asn1.Marshal(*a)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:42
		_go_fuzz_dep_.CoverTab[85962]++
														return edb, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:43
		// _ = "end of CoverTab[85962]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:44
		_go_fuzz_dep_.CoverTab[85963]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:44
		// _ = "end of CoverTab[85963]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:44
	// _ = "end of CoverTab[85960]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:44
	_go_fuzz_dep_.CoverTab[85961]++
													return edb, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:45
	// _ = "end of CoverTab[85961]"
}

// Unmarshal bytes into the EncryptionKey.
func (a *EncryptionKey) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:49
	_go_fuzz_dep_.CoverTab[85964]++
													_, err := asn1.Unmarshal(b, a)
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:51
	// _ = "end of CoverTab[85964]"
}

// Unmarshal bytes into the Checksum.
func (a *Checksum) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:55
	_go_fuzz_dep_.CoverTab[85965]++
													_, err := asn1.Unmarshal(b, a)
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:57
	// _ = "end of CoverTab[85965]"
}

// GenerateEncryptionKey creates a new EncryptionKey with a random key value.
func GenerateEncryptionKey(etype etype.EType) (EncryptionKey, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:61
	_go_fuzz_dep_.CoverTab[85966]++
													k := EncryptionKey{
		KeyType: etype.GetETypeID(),
	}
	b := make([]byte, etype.GetKeyByteSize(), etype.GetKeyByteSize())
	_, err := rand.Read(b)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:67
		_go_fuzz_dep_.CoverTab[85968]++
														return k, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:68
		// _ = "end of CoverTab[85968]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:69
		_go_fuzz_dep_.CoverTab[85969]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:69
		// _ = "end of CoverTab[85969]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:69
	// _ = "end of CoverTab[85966]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:69
	_go_fuzz_dep_.CoverTab[85967]++
													k.KeyValue = b
													return k, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:71
	// _ = "end of CoverTab[85967]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:72
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/Cryptosystem.go:72
var _ = _go_fuzz_dep_.CoverTab
