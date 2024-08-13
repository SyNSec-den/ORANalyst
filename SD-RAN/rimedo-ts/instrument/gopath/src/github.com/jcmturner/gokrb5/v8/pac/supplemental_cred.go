//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:1
package pac

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:1
)

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/jcmturner/rpc/v2/mstypes"
	"github.com/jcmturner/rpc/v2/ndr"
)

const (
	// NTLMSupCredLMOWF indicates that the LM OWF member is present and valid.
	NTLMSupCredLMOWF	uint32	= 31
	// NTLMSupCredNTOWF indicates that the NT OWF member is present and valid.
	NTLMSupCredNTOWF	uint32	= 30
)

// NTLMSupplementalCred implements https://msdn.microsoft.com/en-us/library/cc237949.aspx
type NTLMSupplementalCred struct {
	Version		uint32	// A 32-bit unsigned integer that defines the credential version.This field MUST be 0x00000000.
	Flags		uint32
	LMPassword	[]byte	// A 16-element array of unsigned 8-bit integers that define the LM OWF. The LMPassword member MUST be ignored if the L flag is not set in the Flags member.
	NTPassword	[]byte	// A 16-element array of unsigned 8-bit integers that define the NT OWF. The NTPassword member MUST be ignored if the N flag is not set in the Flags member.
}

// Unmarshal converts the bytes provided into a NTLMSupplementalCred.
func (c *NTLMSupplementalCred) Unmarshal(b []byte) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:29
	_go_fuzz_dep_.CoverTab[87610]++
													r := mstypes.NewReader(bytes.NewReader(b))
													c.Version, err = r.Uint32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:32
		_go_fuzz_dep_.CoverTab[87616]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:33
		// _ = "end of CoverTab[87616]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:34
		_go_fuzz_dep_.CoverTab[87617]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:34
		// _ = "end of CoverTab[87617]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:34
	// _ = "end of CoverTab[87610]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:34
	_go_fuzz_dep_.CoverTab[87611]++
													if c.Version != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:35
		_go_fuzz_dep_.CoverTab[87618]++
														err = errors.New("NTLMSupplementalCred version is not zero")
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:37
		// _ = "end of CoverTab[87618]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:38
		_go_fuzz_dep_.CoverTab[87619]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:38
		// _ = "end of CoverTab[87619]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:38
	// _ = "end of CoverTab[87611]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:38
	_go_fuzz_dep_.CoverTab[87612]++
													c.Flags, err = r.Uint32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:40
		_go_fuzz_dep_.CoverTab[87620]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:41
		// _ = "end of CoverTab[87620]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:42
		_go_fuzz_dep_.CoverTab[87621]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:42
		// _ = "end of CoverTab[87621]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:42
	// _ = "end of CoverTab[87612]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:42
	_go_fuzz_dep_.CoverTab[87613]++
													if isFlagSet(c.Flags, NTLMSupCredLMOWF) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:43
		_go_fuzz_dep_.CoverTab[87622]++
														c.LMPassword, err = r.ReadBytes(16)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:45
			_go_fuzz_dep_.CoverTab[87623]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:46
			// _ = "end of CoverTab[87623]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:47
			_go_fuzz_dep_.CoverTab[87624]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:47
			// _ = "end of CoverTab[87624]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:47
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:47
		// _ = "end of CoverTab[87622]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:48
		_go_fuzz_dep_.CoverTab[87625]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:48
		// _ = "end of CoverTab[87625]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:48
	// _ = "end of CoverTab[87613]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:48
	_go_fuzz_dep_.CoverTab[87614]++
													if isFlagSet(c.Flags, NTLMSupCredNTOWF) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:49
		_go_fuzz_dep_.CoverTab[87626]++
														c.NTPassword, err = r.ReadBytes(16)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:51
			_go_fuzz_dep_.CoverTab[87627]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:52
			// _ = "end of CoverTab[87627]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:53
			_go_fuzz_dep_.CoverTab[87628]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:53
			// _ = "end of CoverTab[87628]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:53
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:53
		// _ = "end of CoverTab[87626]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:54
		_go_fuzz_dep_.CoverTab[87629]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:54
		// _ = "end of CoverTab[87629]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:54
	// _ = "end of CoverTab[87614]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:54
	_go_fuzz_dep_.CoverTab[87615]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:55
	// _ = "end of CoverTab[87615]"
}

// isFlagSet tests if a flag is set in the uint32 little endian flag
func isFlagSet(f uint32, i uint32) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:59
	_go_fuzz_dep_.CoverTab[87630]++

													b := int(i / 8)

													p := uint(7 - (int(i) - 8*b))
													fb := make([]byte, 4)
													binary.LittleEndian.PutUint32(fb, f)
													if fb[b]&(1<<p) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:66
		_go_fuzz_dep_.CoverTab[87632]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:67
		// _ = "end of CoverTab[87632]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:68
		_go_fuzz_dep_.CoverTab[87633]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:68
		// _ = "end of CoverTab[87633]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:68
	// _ = "end of CoverTab[87630]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:68
	_go_fuzz_dep_.CoverTab[87631]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:69
	// _ = "end of CoverTab[87631]"
}

// SECPKGSupplementalCred implements https://msdn.microsoft.com/en-us/library/cc237956.aspx
type SECPKGSupplementalCred struct {
	PackageName	mstypes.RPCUnicodeString
	CredentialSize	uint32
	Credentials	[]uint8	`ndr:"pointer,conformant"`	// Is a ptr. Size is the value of CredentialSize
}

// Unmarshal converts the bytes provided into a SECPKGSupplementalCred.
func (c *SECPKGSupplementalCred) Unmarshal(b []byte) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:80
	_go_fuzz_dep_.CoverTab[87634]++
													dec := ndr.NewDecoder(bytes.NewReader(b))
													err = dec.Decode(c)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:83
		_go_fuzz_dep_.CoverTab[87636]++
														err = fmt.Errorf("error unmarshaling SECPKGSupplementalCred: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:84
		// _ = "end of CoverTab[87636]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:85
		_go_fuzz_dep_.CoverTab[87637]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:85
		// _ = "end of CoverTab[87637]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:85
	// _ = "end of CoverTab[87634]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:85
	_go_fuzz_dep_.CoverTab[87635]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:86
	// _ = "end of CoverTab[87635]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:87
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/supplemental_cred.go:87
var _ = _go_fuzz_dep_.CoverTab
