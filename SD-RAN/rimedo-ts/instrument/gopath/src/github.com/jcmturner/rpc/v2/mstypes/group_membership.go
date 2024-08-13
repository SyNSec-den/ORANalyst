//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:1
package mstypes

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:1
)

// GroupMembership implements https://msdn.microsoft.com/en-us/library/cc237945.aspx
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:3
// RelativeID : A 32-bit unsigned integer that contains the RID of a particular group.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:3
// The possible values for the Attributes flags are identical to those specified in KERB_SID_AND_ATTRIBUTES
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:6
type GroupMembership struct {
	RelativeID	uint32
	Attributes	uint32
}

// DomainGroupMembership implements https://msdn.microsoft.com/en-us/library/hh536344.aspx
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:11
// DomainId: A SID structure that contains the SID for the domain.This member is used in conjunction with the GroupIds members to create group SIDs for the device.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:11
// GroupCount: A 32-bit unsigned integer that contains the number of groups within the domain to which the account belongs.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:11
// GroupIds: A pointer to a list of GROUP_MEMBERSHIP structures that contain the groups to which the account belongs in the domain. The number of groups in this list MUST be equal to GroupCount.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:15
type DomainGroupMembership struct {
	DomainID	RPCSID	`ndr:"pointer"`
	GroupCount	uint32
	GroupIDs	[]GroupMembership	`ndr:"pointer,conformant"`	// Size is value of GroupCount
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:19
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/group_membership.go:19
var _ = _go_fuzz_dep_.CoverTab
