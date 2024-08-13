//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:1
// Package pac implements Microsoft Privilege Attribute Certificate (PAC) processing.
package pac

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:2
)

import (
	"bytes"
	"fmt"

	"github.com/jcmturner/rpc/v2/mstypes"
	"github.com/jcmturner/rpc/v2/ndr"
)

// KERB_VALIDATION_INFO flags.
const (
	USERFLAG_GUEST						= 31	// Authentication was done via the GUEST account; no password was used.
	USERFLAG_NO_ENCRYPTION_AVAILABLE			= 30	// No encryption is available.
	USERFLAG_LAN_MANAGER_KEY				= 28	// LAN Manager key was used for authentication.
	USERFLAG_SUB_AUTH					= 25	// Sub-authentication used; session key came from the sub-authentication package.
	USERFLAG_EXTRA_SIDS					= 26	// Indicates that the ExtraSids field is populated and contains additional SIDs.
	USERFLAG_MACHINE_ACCOUNT				= 24	// Indicates that the account is a machine account.
	USERFLAG_DC_NTLM2					= 23	// Indicates that the domain controller understands NTLMv2.
	USERFLAG_RESOURCE_GROUPIDS				= 22	// Indicates that the ResourceGroupIds field is populated.
	USERFLAG_PROFILEPATH					= 21	// Indicates that ProfilePath is populated.
	USERFLAG_NTLM2_NTCHALLENGERESP				= 20	// The NTLMv2 response from the NtChallengeResponseFields ([MS-NLMP] section 2.2.1.3) was used for authentication and session key generation.
	USERFLAG_LM2_LMCHALLENGERESP				= 19	// The LMv2 response from the LmChallengeResponseFields ([MS-NLMP] section 2.2.1.3) was used for authentication and session key generation.
	USERFLAG_AUTH_LMCHALLENGERESP_KEY_NTCHALLENGERESP	= 18	// The LMv2 response from the LmChallengeResponseFields ([MS-NLMP] section 2.2.1.3) was used for authentication and the NTLMv2 response from the NtChallengeResponseFields ([MS-NLMP] section 2.2.1.3) was used session key generation.
)

// KerbValidationInfo implement https://msdn.microsoft.com/en-us/library/cc237948.aspx
type KerbValidationInfo struct {
	LogOnTime		mstypes.FileTime
	LogOffTime		mstypes.FileTime
	KickOffTime		mstypes.FileTime
	PasswordLastSet		mstypes.FileTime
	PasswordCanChange	mstypes.FileTime
	PasswordMustChange	mstypes.FileTime
	EffectiveName		mstypes.RPCUnicodeString
	FullName		mstypes.RPCUnicodeString
	LogonScript		mstypes.RPCUnicodeString
	ProfilePath		mstypes.RPCUnicodeString
	HomeDirectory		mstypes.RPCUnicodeString
	HomeDirectoryDrive	mstypes.RPCUnicodeString
	LogonCount		uint16
	BadPasswordCount	uint16
	UserID			uint32
	PrimaryGroupID		uint32
	GroupCount		uint32
	GroupIDs		[]mstypes.GroupMembership	`ndr:"pointer,conformant"`
	UserFlags		uint32
	UserSessionKey		mstypes.UserSessionKey
	LogonServer		mstypes.RPCUnicodeString
	LogonDomainName		mstypes.RPCUnicodeString
	LogonDomainID		mstypes.RPCSID	`ndr:"pointer"`
	Reserved1		[2]uint32	// Has 2 elements
	UserAccountControl	uint32
	SubAuthStatus		uint32
	LastSuccessfulILogon	mstypes.FileTime
	LastFailedILogon	mstypes.FileTime
	FailedILogonCount	uint32
	Reserved3		uint32
	SIDCount		uint32
	ExtraSIDs		[]mstypes.KerbSidAndAttributes	`ndr:"pointer,conformant"`
	ResourceGroupDomainSID	mstypes.RPCSID			`ndr:"pointer"`
	ResourceGroupCount	uint32
	ResourceGroupIDs	[]mstypes.GroupMembership	`ndr:"pointer,conformant"`
}

// Unmarshal bytes into the DeviceInfo struct
func (k *KerbValidationInfo) Unmarshal(b []byte) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:68
	_go_fuzz_dep_.CoverTab[87458]++
													dec := ndr.NewDecoder(bytes.NewReader(b))
													err = dec.Decode(k)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:71
		_go_fuzz_dep_.CoverTab[87460]++
														err = fmt.Errorf("error unmarshaling KerbValidationInfo: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:72
		// _ = "end of CoverTab[87460]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:73
		_go_fuzz_dep_.CoverTab[87461]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:73
		// _ = "end of CoverTab[87461]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:73
	// _ = "end of CoverTab[87458]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:73
	_go_fuzz_dep_.CoverTab[87459]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:74
	// _ = "end of CoverTab[87459]"
}

// GetGroupMembershipSIDs returns a slice of strings containing the group membership SIDs found in the PAC.
func (k *KerbValidationInfo) GetGroupMembershipSIDs() []string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:78
	_go_fuzz_dep_.CoverTab[87462]++
													var g []string
													lSID := k.LogonDomainID.String()
													for i := range k.GroupIDs {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:81
		_go_fuzz_dep_.CoverTab[87466]++
														g = append(g, fmt.Sprintf("%s-%d", lSID, k.GroupIDs[i].RelativeID))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:82
		// _ = "end of CoverTab[87466]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:83
	// _ = "end of CoverTab[87462]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:83
	_go_fuzz_dep_.CoverTab[87463]++
													for _, s := range k.ExtraSIDs {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:84
		_go_fuzz_dep_.CoverTab[87467]++
														var exists = false
														for _, es := range g {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:86
			_go_fuzz_dep_.CoverTab[87469]++
															if es == s.SID.String() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:87
				_go_fuzz_dep_.CoverTab[87470]++
																exists = true
																break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:89
				// _ = "end of CoverTab[87470]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:90
				_go_fuzz_dep_.CoverTab[87471]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:90
				// _ = "end of CoverTab[87471]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:90
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:90
			// _ = "end of CoverTab[87469]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:91
		// _ = "end of CoverTab[87467]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:91
		_go_fuzz_dep_.CoverTab[87468]++
														if !exists {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:92
			_go_fuzz_dep_.CoverTab[87472]++
															g = append(g, s.SID.String())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:93
			// _ = "end of CoverTab[87472]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:94
			_go_fuzz_dep_.CoverTab[87473]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:94
			// _ = "end of CoverTab[87473]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:94
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:94
		// _ = "end of CoverTab[87468]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:95
	// _ = "end of CoverTab[87463]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:95
	_go_fuzz_dep_.CoverTab[87464]++
													for _, r := range k.ResourceGroupIDs {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:96
		_go_fuzz_dep_.CoverTab[87474]++
														var exists = false
														s := fmt.Sprintf("%s-%d", k.ResourceGroupDomainSID.String(), r.RelativeID)
														for _, es := range g {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:99
			_go_fuzz_dep_.CoverTab[87476]++
															if es == s {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:100
				_go_fuzz_dep_.CoverTab[87477]++
																	exists = true
																	break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:102
				// _ = "end of CoverTab[87477]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:103
				_go_fuzz_dep_.CoverTab[87478]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:103
				// _ = "end of CoverTab[87478]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:103
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:103
			// _ = "end of CoverTab[87476]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:104
		// _ = "end of CoverTab[87474]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:104
		_go_fuzz_dep_.CoverTab[87475]++
															if !exists {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:105
			_go_fuzz_dep_.CoverTab[87479]++
																g = append(g, s)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:106
			// _ = "end of CoverTab[87479]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:107
			_go_fuzz_dep_.CoverTab[87480]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:107
			// _ = "end of CoverTab[87480]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:107
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:107
		// _ = "end of CoverTab[87475]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:108
	// _ = "end of CoverTab[87464]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:108
	_go_fuzz_dep_.CoverTab[87465]++
														return g
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:109
	// _ = "end of CoverTab[87465]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:110
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/kerb_validation_info.go:110
var _ = _go_fuzz_dep_.CoverTab
