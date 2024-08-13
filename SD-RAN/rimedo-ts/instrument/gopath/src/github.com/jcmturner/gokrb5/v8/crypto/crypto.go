//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:1
// Package crypto implements cryptographic functions for Kerberos 5 implementation.
package crypto

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:2
)

import (
	"encoding/hex"
	"fmt"

	"github.com/jcmturner/gokrb5/v8/crypto/etype"
	"github.com/jcmturner/gokrb5/v8/iana/chksumtype"
	"github.com/jcmturner/gokrb5/v8/iana/etypeID"
	"github.com/jcmturner/gokrb5/v8/iana/patype"
	"github.com/jcmturner/gokrb5/v8/types"
)

// GetEtype returns an instances of the required etype struct for the etype ID.
func GetEtype(id int32) (etype.EType, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:16
	_go_fuzz_dep_.CoverTab[86199]++
												switch id {
	case etypeID.AES128_CTS_HMAC_SHA1_96:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:18
		_go_fuzz_dep_.CoverTab[86200]++
													var et Aes128CtsHmacSha96
													return et, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:20
		// _ = "end of CoverTab[86200]"
	case etypeID.AES256_CTS_HMAC_SHA1_96:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:21
		_go_fuzz_dep_.CoverTab[86201]++
													var et Aes256CtsHmacSha96
													return et, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:23
		// _ = "end of CoverTab[86201]"
	case etypeID.AES128_CTS_HMAC_SHA256_128:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:24
		_go_fuzz_dep_.CoverTab[86202]++
													var et Aes128CtsHmacSha256128
													return et, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:26
		// _ = "end of CoverTab[86202]"
	case etypeID.AES256_CTS_HMAC_SHA384_192:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:27
		_go_fuzz_dep_.CoverTab[86203]++
													var et Aes256CtsHmacSha384192
													return et, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:29
		// _ = "end of CoverTab[86203]"
	case etypeID.DES3_CBC_SHA1_KD:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:30
		_go_fuzz_dep_.CoverTab[86204]++
													var et Des3CbcSha1Kd
													return et, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:32
		// _ = "end of CoverTab[86204]"
	case etypeID.RC4_HMAC:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:33
		_go_fuzz_dep_.CoverTab[86205]++
													var et RC4HMAC
													return et, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:35
		// _ = "end of CoverTab[86205]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:36
		_go_fuzz_dep_.CoverTab[86206]++
													return nil, fmt.Errorf("unknown or unsupported EType: %d", id)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:37
		// _ = "end of CoverTab[86206]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:38
	// _ = "end of CoverTab[86199]"
}

// GetChksumEtype returns an instances of the required etype struct for the checksum ID.
func GetChksumEtype(id int32) (etype.EType, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:42
	_go_fuzz_dep_.CoverTab[86207]++
												switch id {
	case chksumtype.HMAC_SHA1_96_AES128:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:44
		_go_fuzz_dep_.CoverTab[86208]++
													var et Aes128CtsHmacSha96
													return et, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:46
		// _ = "end of CoverTab[86208]"
	case chksumtype.HMAC_SHA1_96_AES256:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:47
		_go_fuzz_dep_.CoverTab[86209]++
													var et Aes256CtsHmacSha96
													return et, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:49
		// _ = "end of CoverTab[86209]"
	case chksumtype.HMAC_SHA256_128_AES128:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:50
		_go_fuzz_dep_.CoverTab[86210]++
													var et Aes128CtsHmacSha256128
													return et, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:52
		// _ = "end of CoverTab[86210]"
	case chksumtype.HMAC_SHA384_192_AES256:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:53
		_go_fuzz_dep_.CoverTab[86211]++
													var et Aes256CtsHmacSha384192
													return et, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:55
		// _ = "end of CoverTab[86211]"
	case chksumtype.HMAC_SHA1_DES3_KD:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:56
		_go_fuzz_dep_.CoverTab[86212]++
													var et Des3CbcSha1Kd
													return et, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:58
		// _ = "end of CoverTab[86212]"
	case chksumtype.KERB_CHECKSUM_HMAC_MD5:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:59
		_go_fuzz_dep_.CoverTab[86213]++
													var et RC4HMAC
													return et, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:61
		// _ = "end of CoverTab[86213]"

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:65
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:65
		_go_fuzz_dep_.CoverTab[86214]++
													return nil, fmt.Errorf("unknown or unsupported checksum type: %d", id)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:66
		// _ = "end of CoverTab[86214]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:67
	// _ = "end of CoverTab[86207]"
}

// GetKeyFromPassword generates an encryption key from the principal's password.
func GetKeyFromPassword(passwd string, cname types.PrincipalName, realm string, etypeID int32, pas types.PADataSequence) (types.EncryptionKey, etype.EType, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:71
	_go_fuzz_dep_.CoverTab[86215]++
												var key types.EncryptionKey
												et, err := GetEtype(etypeID)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:74
		_go_fuzz_dep_.CoverTab[86220]++
													return key, et, fmt.Errorf("error getting encryption type: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:75
		// _ = "end of CoverTab[86220]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:76
		_go_fuzz_dep_.CoverTab[86221]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:76
		// _ = "end of CoverTab[86221]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:76
	// _ = "end of CoverTab[86215]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:76
	_go_fuzz_dep_.CoverTab[86216]++
												sk2p := et.GetDefaultStringToKeyParams()
												var salt string
												var paID int32
												for _, pa := range pas {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:80
		_go_fuzz_dep_.CoverTab[86222]++
													switch pa.PADataType {
		case patype.PA_PW_SALT:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:82
			_go_fuzz_dep_.CoverTab[86223]++
														if paID > pa.PADataType {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:83
				_go_fuzz_dep_.CoverTab[86235]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:84
				// _ = "end of CoverTab[86235]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:85
				_go_fuzz_dep_.CoverTab[86236]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:85
				// _ = "end of CoverTab[86236]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:85
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:85
			// _ = "end of CoverTab[86223]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:85
			_go_fuzz_dep_.CoverTab[86224]++
														salt = string(pa.PADataValue)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:86
			// _ = "end of CoverTab[86224]"
		case patype.PA_ETYPE_INFO:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:87
			_go_fuzz_dep_.CoverTab[86225]++
														if paID > pa.PADataType {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:88
				_go_fuzz_dep_.CoverTab[86237]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:89
				// _ = "end of CoverTab[86237]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:90
				_go_fuzz_dep_.CoverTab[86238]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:90
				// _ = "end of CoverTab[86238]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:90
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:90
			// _ = "end of CoverTab[86225]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:90
			_go_fuzz_dep_.CoverTab[86226]++
														var eti types.ETypeInfo
														err := eti.Unmarshal(pa.PADataValue)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:93
				_go_fuzz_dep_.CoverTab[86239]++
															return key, et, fmt.Errorf("error unmashaling PA Data to PA-ETYPE-INFO2: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:94
				// _ = "end of CoverTab[86239]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:95
				_go_fuzz_dep_.CoverTab[86240]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:95
				// _ = "end of CoverTab[86240]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:95
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:95
			// _ = "end of CoverTab[86226]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:95
			_go_fuzz_dep_.CoverTab[86227]++
														if etypeID != eti[0].EType {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:96
				_go_fuzz_dep_.CoverTab[86241]++
															et, err = GetEtype(eti[0].EType)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:98
					_go_fuzz_dep_.CoverTab[86242]++
																return key, et, fmt.Errorf("error getting encryption type: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:99
					// _ = "end of CoverTab[86242]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:100
					_go_fuzz_dep_.CoverTab[86243]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:100
					// _ = "end of CoverTab[86243]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:100
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:100
				// _ = "end of CoverTab[86241]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:101
				_go_fuzz_dep_.CoverTab[86244]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:101
				// _ = "end of CoverTab[86244]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:101
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:101
			// _ = "end of CoverTab[86227]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:101
			_go_fuzz_dep_.CoverTab[86228]++
														salt = string(eti[0].Salt)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:102
			// _ = "end of CoverTab[86228]"
		case patype.PA_ETYPE_INFO2:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:103
			_go_fuzz_dep_.CoverTab[86229]++
														if paID > pa.PADataType {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:104
				_go_fuzz_dep_.CoverTab[86245]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:105
				// _ = "end of CoverTab[86245]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:106
				_go_fuzz_dep_.CoverTab[86246]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:106
				// _ = "end of CoverTab[86246]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:106
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:106
			// _ = "end of CoverTab[86229]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:106
			_go_fuzz_dep_.CoverTab[86230]++
														var et2 types.ETypeInfo2
														err := et2.Unmarshal(pa.PADataValue)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:109
				_go_fuzz_dep_.CoverTab[86247]++
															return key, et, fmt.Errorf("error unmashalling PA Data to PA-ETYPE-INFO2: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:110
				// _ = "end of CoverTab[86247]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:111
				_go_fuzz_dep_.CoverTab[86248]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:111
				// _ = "end of CoverTab[86248]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:111
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:111
			// _ = "end of CoverTab[86230]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:111
			_go_fuzz_dep_.CoverTab[86231]++
														if etypeID != et2[0].EType {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:112
				_go_fuzz_dep_.CoverTab[86249]++
															et, err = GetEtype(et2[0].EType)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:114
					_go_fuzz_dep_.CoverTab[86250]++
																return key, et, fmt.Errorf("error getting encryption type: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:115
					// _ = "end of CoverTab[86250]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:116
					_go_fuzz_dep_.CoverTab[86251]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:116
					// _ = "end of CoverTab[86251]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:116
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:116
				// _ = "end of CoverTab[86249]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:117
				_go_fuzz_dep_.CoverTab[86252]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:117
				// _ = "end of CoverTab[86252]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:117
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:117
			// _ = "end of CoverTab[86231]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:117
			_go_fuzz_dep_.CoverTab[86232]++
														if len(et2[0].S2KParams) == 4 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:118
				_go_fuzz_dep_.CoverTab[86253]++
															sk2p = hex.EncodeToString(et2[0].S2KParams)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:119
				// _ = "end of CoverTab[86253]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:120
				_go_fuzz_dep_.CoverTab[86254]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:120
				// _ = "end of CoverTab[86254]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:120
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:120
			// _ = "end of CoverTab[86232]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:120
			_go_fuzz_dep_.CoverTab[86233]++
														salt = et2[0].Salt
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:121
			// _ = "end of CoverTab[86233]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:121
		default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:121
			_go_fuzz_dep_.CoverTab[86234]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:121
			// _ = "end of CoverTab[86234]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:122
		// _ = "end of CoverTab[86222]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:123
	// _ = "end of CoverTab[86216]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:123
	_go_fuzz_dep_.CoverTab[86217]++
												if salt == "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:124
		_go_fuzz_dep_.CoverTab[86255]++
													salt = cname.GetSalt(realm)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:125
		// _ = "end of CoverTab[86255]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:126
		_go_fuzz_dep_.CoverTab[86256]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:126
		// _ = "end of CoverTab[86256]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:126
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:126
	// _ = "end of CoverTab[86217]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:126
	_go_fuzz_dep_.CoverTab[86218]++
												k, err := et.StringToKey(passwd, salt, sk2p)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:128
		_go_fuzz_dep_.CoverTab[86257]++
													return key, et, fmt.Errorf("error deriving key from string: %+v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:129
		// _ = "end of CoverTab[86257]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:130
		_go_fuzz_dep_.CoverTab[86258]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:130
		// _ = "end of CoverTab[86258]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:130
	// _ = "end of CoverTab[86218]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:130
	_go_fuzz_dep_.CoverTab[86219]++
												key = types.EncryptionKey{
		KeyType:	etypeID,
		KeyValue:	k,
	}
												return key, et, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:135
	// _ = "end of CoverTab[86219]"
}

// GetEncryptedData encrypts the data provided and returns and EncryptedData type.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:138
// Pass a usage value of zero to use the key provided directly rather than deriving one.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:140
func GetEncryptedData(plainBytes []byte, key types.EncryptionKey, usage uint32, kvno int) (types.EncryptedData, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:140
	_go_fuzz_dep_.CoverTab[86259]++
												var ed types.EncryptedData
												et, err := GetEtype(key.KeyType)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:143
		_go_fuzz_dep_.CoverTab[86262]++
													return ed, fmt.Errorf("error getting etype: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:144
		// _ = "end of CoverTab[86262]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:145
		_go_fuzz_dep_.CoverTab[86263]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:145
		// _ = "end of CoverTab[86263]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:145
	// _ = "end of CoverTab[86259]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:145
	_go_fuzz_dep_.CoverTab[86260]++
												_, b, err := et.EncryptMessage(key.KeyValue, plainBytes, usage)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:147
		_go_fuzz_dep_.CoverTab[86264]++
													return ed, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:148
		// _ = "end of CoverTab[86264]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:149
		_go_fuzz_dep_.CoverTab[86265]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:149
		// _ = "end of CoverTab[86265]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:149
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:149
	// _ = "end of CoverTab[86260]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:149
	_go_fuzz_dep_.CoverTab[86261]++

												ed = types.EncryptedData{
		EType:	key.KeyType,
		Cipher:	b,
		KVNO:	kvno,
	}
												return ed, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:156
	// _ = "end of CoverTab[86261]"
}

// DecryptEncPart decrypts the EncryptedData.
func DecryptEncPart(ed types.EncryptedData, key types.EncryptionKey, usage uint32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:160
	_go_fuzz_dep_.CoverTab[86266]++
												return DecryptMessage(ed.Cipher, key, usage)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:161
	// _ = "end of CoverTab[86266]"
}

// DecryptMessage decrypts the ciphertext and verifies the integrity.
func DecryptMessage(ciphertext []byte, key types.EncryptionKey, usage uint32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:165
	_go_fuzz_dep_.CoverTab[86267]++
												et, err := GetEtype(key.KeyType)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:167
		_go_fuzz_dep_.CoverTab[86270]++
													return []byte{}, fmt.Errorf("error decrypting: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:168
		// _ = "end of CoverTab[86270]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:169
		_go_fuzz_dep_.CoverTab[86271]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:169
		// _ = "end of CoverTab[86271]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:169
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:169
	// _ = "end of CoverTab[86267]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:169
	_go_fuzz_dep_.CoverTab[86268]++
												b, err := et.DecryptMessage(key.KeyValue, ciphertext, usage)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:171
		_go_fuzz_dep_.CoverTab[86272]++
													return nil, fmt.Errorf("error decrypting: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:172
		// _ = "end of CoverTab[86272]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:173
		_go_fuzz_dep_.CoverTab[86273]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:173
		// _ = "end of CoverTab[86273]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:173
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:173
	// _ = "end of CoverTab[86268]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:173
	_go_fuzz_dep_.CoverTab[86269]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:174
	// _ = "end of CoverTab[86269]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:175
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/crypto.go:175
var _ = _go_fuzz_dep_.CoverTab
