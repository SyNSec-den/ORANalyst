//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:1
package pac

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:1
)

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/types"
	"github.com/jcmturner/rpc/v2/mstypes"
	"github.com/jcmturner/rpc/v2/ndr"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:17
// CredentialsInfo implements https://msdn.microsoft.com/en-us/library/cc237953.aspx
type CredentialsInfo struct {
	Version				uint32	// A 32-bit unsigned integer in little-endian format that defines the version. MUST be 0x00000000.
	EType				uint32
	PACCredentialDataEncrypted	[]byte	// Key usage number for encryption: KERB_NON_KERB_SALT (16)
	PACCredentialData		CredentialData
}

// Unmarshal bytes into the CredentialsInfo struct
func (c *CredentialsInfo) Unmarshal(b []byte, k types.EncryptionKey) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:26
	_go_fuzz_dep_.CoverTab[87417]++

													r := mstypes.NewReader(bytes.NewReader(b))

													c.Version, err = r.Uint32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:31
		_go_fuzz_dep_.CoverTab[87423]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:32
		// _ = "end of CoverTab[87423]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:33
		_go_fuzz_dep_.CoverTab[87424]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:33
		// _ = "end of CoverTab[87424]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:33
	// _ = "end of CoverTab[87417]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:33
	_go_fuzz_dep_.CoverTab[87418]++
													if c.Version != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:34
		_go_fuzz_dep_.CoverTab[87425]++
														err = errors.New("credentials info version is not zero")
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:36
		// _ = "end of CoverTab[87425]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:37
		_go_fuzz_dep_.CoverTab[87426]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:37
		// _ = "end of CoverTab[87426]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:37
	// _ = "end of CoverTab[87418]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:37
	_go_fuzz_dep_.CoverTab[87419]++
													c.EType, err = r.Uint32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:39
		_go_fuzz_dep_.CoverTab[87427]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:40
		// _ = "end of CoverTab[87427]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:41
		_go_fuzz_dep_.CoverTab[87428]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:41
		// _ = "end of CoverTab[87428]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:41
	// _ = "end of CoverTab[87419]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:41
	_go_fuzz_dep_.CoverTab[87420]++
													c.PACCredentialDataEncrypted, err = r.ReadBytes(len(b) - 8)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:43
		_go_fuzz_dep_.CoverTab[87429]++
														err = fmt.Errorf("error reading PAC Credetials Data: %v", err)
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:45
		// _ = "end of CoverTab[87429]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:46
		_go_fuzz_dep_.CoverTab[87430]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:46
		// _ = "end of CoverTab[87430]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:46
	// _ = "end of CoverTab[87420]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:46
	_go_fuzz_dep_.CoverTab[87421]++

													err = c.DecryptEncPart(k)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:49
		_go_fuzz_dep_.CoverTab[87431]++
														err = fmt.Errorf("error decrypting PAC Credentials Data: %v", err)
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:51
		// _ = "end of CoverTab[87431]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:52
		_go_fuzz_dep_.CoverTab[87432]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:52
		// _ = "end of CoverTab[87432]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:52
	// _ = "end of CoverTab[87421]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:52
	_go_fuzz_dep_.CoverTab[87422]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:53
	// _ = "end of CoverTab[87422]"
}

// DecryptEncPart decrypts the encrypted part of the CredentialsInfo.
func (c *CredentialsInfo) DecryptEncPart(k types.EncryptionKey) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:57
	_go_fuzz_dep_.CoverTab[87433]++
													if k.KeyType != int32(c.EType) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:58
		_go_fuzz_dep_.CoverTab[87437]++
														return fmt.Errorf("key provided is not the correct type. Type needed: %d, type provided: %d", c.EType, k.KeyType)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:59
		// _ = "end of CoverTab[87437]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:60
		_go_fuzz_dep_.CoverTab[87438]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:60
		// _ = "end of CoverTab[87438]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:60
	// _ = "end of CoverTab[87433]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:60
	_go_fuzz_dep_.CoverTab[87434]++
													pt, err := crypto.DecryptMessage(c.PACCredentialDataEncrypted, k, keyusage.KERB_NON_KERB_SALT)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:62
		_go_fuzz_dep_.CoverTab[87439]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:63
		// _ = "end of CoverTab[87439]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:64
		_go_fuzz_dep_.CoverTab[87440]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:64
		// _ = "end of CoverTab[87440]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:64
	// _ = "end of CoverTab[87434]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:64
	_go_fuzz_dep_.CoverTab[87435]++
													err = c.PACCredentialData.Unmarshal(pt)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:66
		_go_fuzz_dep_.CoverTab[87441]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:67
		// _ = "end of CoverTab[87441]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:68
		_go_fuzz_dep_.CoverTab[87442]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:68
		// _ = "end of CoverTab[87442]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:68
	// _ = "end of CoverTab[87435]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:68
	_go_fuzz_dep_.CoverTab[87436]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:69
	// _ = "end of CoverTab[87436]"
}

// CredentialData implements https://msdn.microsoft.com/en-us/library/cc237952.aspx
type CredentialData struct {
	CredentialCount	uint32
	Credentials	[]SECPKGSupplementalCred	// Size is the value of CredentialCount
}

// Unmarshal converts the bytes provided into a CredentialData type.
func (c *CredentialData) Unmarshal(b []byte) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:79
	_go_fuzz_dep_.CoverTab[87443]++
													dec := ndr.NewDecoder(bytes.NewReader(b))
													err = dec.Decode(c)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:82
		_go_fuzz_dep_.CoverTab[87445]++
														err = fmt.Errorf("error unmarshaling KerbValidationInfo: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:83
		// _ = "end of CoverTab[87445]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:84
		_go_fuzz_dep_.CoverTab[87446]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:84
		// _ = "end of CoverTab[87446]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:84
	// _ = "end of CoverTab[87443]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:84
	_go_fuzz_dep_.CoverTab[87444]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:85
	// _ = "end of CoverTab[87444]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:86
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/credentials_info.go:86
var _ = _go_fuzz_dep_.CoverTab
