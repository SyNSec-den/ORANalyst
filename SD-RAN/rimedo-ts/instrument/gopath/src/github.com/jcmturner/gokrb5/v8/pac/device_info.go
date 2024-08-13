//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:1
package pac

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:1
)

import (
	"bytes"
	"fmt"

	"github.com/jcmturner/rpc/v2/mstypes"
	"github.com/jcmturner/rpc/v2/ndr"
)

// DeviceInfo implements https://msdn.microsoft.com/en-us/library/hh536402.aspx
type DeviceInfo struct {
	UserID			uint32				// A 32-bit unsigned integer that contains the RID of the account. If the UserId member equals 0x00000000, the first group SID in this member is the SID for this account.
	PrimaryGroupID		uint32				// A 32-bit unsigned integer that contains the RID for the primary group to which this account belongs.
	AccountDomainID		mstypes.RPCSID			`ndr:"pointer"`	// A SID structure that contains the SID for the domain of the account.This member is used in conjunction with the UserId, and GroupIds members to create the user and group SIDs for the client.
	AccountGroupCount	uint32				// A 32-bit unsigned integer that contains the number of groups within the account domain to which the account belongs
	AccountGroupIDs		[]mstypes.GroupMembership	`ndr:"pointer,conformant"`	// A pointer to a list of GROUP_MEMBERSHIP (section 2.2.2) structures that contains the groups to which the account belongs in the account domain. The number of groups in this list MUST be equal to GroupCount.
	SIDCount		uint32				// A 32-bit unsigned integer that contains the total number of SIDs present in the ExtraSids member.
	ExtraSIDs		[]mstypes.KerbSidAndAttributes	`ndr:"pointer,conformant"`	// A pointer to a list of KERB_SID_AND_ATTRIBUTES structures that contain a list of SIDs corresponding to groups not in domains. If the UserId member equals 0x00000000, the first group SID in this member is the SID for this account.
	DomainGroupCount	uint32				// A 32-bit unsigned integer that contains the number of domains with groups to which the account belongs.
	DomainGroup		[]mstypes.DomainGroupMembership	`ndr:"pointer,conformant"`	// A pointer to a list of DOMAIN_GROUP_MEMBERSHIP structures (section 2.2.3) that contains the domains to which the account belongs to a group. The number of sets in this list MUST be equal to DomainCount.
}

// Unmarshal bytes into the DeviceInfo struct
func (k *DeviceInfo) Unmarshal(b []byte) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:25
	_go_fuzz_dep_.CoverTab[87454]++
												dec := ndr.NewDecoder(bytes.NewReader(b))
												err = dec.Decode(k)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:28
		_go_fuzz_dep_.CoverTab[87456]++
													err = fmt.Errorf("error unmarshaling DeviceInfo: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:29
		// _ = "end of CoverTab[87456]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:30
		_go_fuzz_dep_.CoverTab[87457]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:30
		// _ = "end of CoverTab[87457]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:30
	// _ = "end of CoverTab[87454]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:30
	_go_fuzz_dep_.CoverTab[87455]++
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:31
	// _ = "end of CoverTab[87455]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:32
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_info.go:32
var _ = _go_fuzz_dep_.CoverTab
