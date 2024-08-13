//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/kerb_sid_and_attributes.go:1
package mstypes

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/kerb_sid_and_attributes.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/kerb_sid_and_attributes.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/kerb_sid_and_attributes.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/kerb_sid_and_attributes.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/kerb_sid_and_attributes.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/kerb_sid_and_attributes.go:1
)

// Attributes of a security group membership and can be combined by using the bitwise OR operation.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/kerb_sid_and_attributes.go:3
// They are used by an access check mechanism to specify whether the membership is to be used in an access check decision.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/kerb_sid_and_attributes.go:5
const (
	SEGroupMandatory	= 31
	SEGroupEnabledByDefault	= 30
	SEGroupEnabled		= 29
	SEGroupOwner		= 28
	SEGroupResource		= 2
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/kerb_sid_and_attributes.go:12
)

// KerbSidAndAttributes implements https://msdn.microsoft.com/en-us/library/cc237947.aspx
type KerbSidAndAttributes struct {
	SID		RPCSID	`ndr:"pointer"`	// A pointer to an RPC_SID structure.
	Attributes	uint32
}

// SetFlag sets a flag in a uint32 attribute value.
func SetFlag(a *uint32, i uint) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/kerb_sid_and_attributes.go:21
	_go_fuzz_dep_.CoverTab[87359]++
														*a = *a | (1 << (31 - i))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/kerb_sid_and_attributes.go:22
	// _ = "end of CoverTab[87359]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/kerb_sid_and_attributes.go:23
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/kerb_sid_and_attributes.go:23
var _ = _go_fuzz_dep_.CoverTab
