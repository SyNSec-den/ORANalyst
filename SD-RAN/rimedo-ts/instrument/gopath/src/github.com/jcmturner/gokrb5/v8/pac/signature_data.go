//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:1
package pac

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:1
)

import (
	"bytes"

	"github.com/jcmturner/gokrb5/v8/iana/chksumtype"
	"github.com/jcmturner/rpc/v2/mstypes"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:22
// SignatureData implements https://msdn.microsoft.com/en-us/library/cc237955.aspx
type SignatureData struct {
	SignatureType	uint32	// A 32-bit unsigned integer value in little-endian format that defines the cryptographic system used to calculate the checksum. This MUST be one of the following checksum types: KERB_CHECKSUM_HMAC_MD5 (signature size = 16), HMAC_SHA1_96_AES128 (signature size = 12), HMAC_SHA1_96_AES256 (signature size = 12).
	Signature	[]byte	// Size depends on the type. See comment above.
	RODCIdentifier	uint16	// A 16-bit unsigned integer value in little-endian format that contains the first 16 bits of the key version number ([MS-KILE] section 3.1.5.8) when the KDC is an RODC. When the KDC is not an RODC, this field does not exist.
}

// Unmarshal bytes into the SignatureData struct
func (k *SignatureData) Unmarshal(b []byte) (rb []byte, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:30
	_go_fuzz_dep_.CoverTab[87593]++
													r := mstypes.NewReader(bytes.NewReader(b))

													k.SignatureType, err = r.Uint32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:34
		_go_fuzz_dep_.CoverTab[87598]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:35
		// _ = "end of CoverTab[87598]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:36
		_go_fuzz_dep_.CoverTab[87599]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:36
		// _ = "end of CoverTab[87599]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:36
	// _ = "end of CoverTab[87593]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:36
	_go_fuzz_dep_.CoverTab[87594]++

													var c int
													switch k.SignatureType {
	case chksumtype.KERB_CHECKSUM_HMAC_MD5_UNSIGNED:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:40
		_go_fuzz_dep_.CoverTab[87600]++
														c = 16
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:41
		// _ = "end of CoverTab[87600]"
	case uint32(chksumtype.HMAC_SHA1_96_AES128):
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:42
		_go_fuzz_dep_.CoverTab[87601]++
														c = 12
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:43
		// _ = "end of CoverTab[87601]"
	case uint32(chksumtype.HMAC_SHA1_96_AES256):
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:44
		_go_fuzz_dep_.CoverTab[87602]++
														c = 12
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:45
		// _ = "end of CoverTab[87602]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:45
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:45
		_go_fuzz_dep_.CoverTab[87603]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:45
		// _ = "end of CoverTab[87603]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:46
	// _ = "end of CoverTab[87594]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:46
	_go_fuzz_dep_.CoverTab[87595]++
													k.Signature, err = r.ReadBytes(c)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:48
		_go_fuzz_dep_.CoverTab[87604]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:49
		// _ = "end of CoverTab[87604]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:50
		_go_fuzz_dep_.CoverTab[87605]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:50
		// _ = "end of CoverTab[87605]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:50
	// _ = "end of CoverTab[87595]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:50
	_go_fuzz_dep_.CoverTab[87596]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:53
	if len(b) >= 4+c+2 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:53
		_go_fuzz_dep_.CoverTab[87606]++
														k.RODCIdentifier, err = r.Uint16()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:55
			_go_fuzz_dep_.CoverTab[87607]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:56
			// _ = "end of CoverTab[87607]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:57
			_go_fuzz_dep_.CoverTab[87608]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:57
			// _ = "end of CoverTab[87608]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:57
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:57
		// _ = "end of CoverTab[87606]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:58
		_go_fuzz_dep_.CoverTab[87609]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:58
		// _ = "end of CoverTab[87609]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:58
	// _ = "end of CoverTab[87596]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:58
	_go_fuzz_dep_.CoverTab[87597]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:61
	rb = make([]byte, len(b), len(b))
													copy(rb, b)
													z := make([]byte, len(b), len(b))
													copy(rb[4:4+c], z)

													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:66
	// _ = "end of CoverTab[87597]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:67
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/signature_data.go:67
var _ = _go_fuzz_dep_.CoverTab
