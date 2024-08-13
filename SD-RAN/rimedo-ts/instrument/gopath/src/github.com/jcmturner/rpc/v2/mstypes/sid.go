//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:1
package mstypes

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:1
)

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

// RPCSID implements https://msdn.microsoft.com/en-us/library/cc230364.aspx
type RPCSID struct {
	Revision		uint8		// An 8-bit unsigned integer that specifies the revision level of the SID. This value MUST be set to 0x01.
	SubAuthorityCount	uint8		// An 8-bit unsigned integer that specifies the number of elements in the SubAuthority array. The maximum number of elements allowed is 15.
	IdentifierAuthority	[6]byte		// An RPC_SID_IDENTIFIER_AUTHORITY structure that indicates the authority under which the SID was created. It describes the entity that created the SID. The Identifier Authority value {0,0,0,0,0,5} denotes SIDs created by the NT SID authority.
	SubAuthority		[]uint32	`ndr:"conformant"`	// A variable length array of unsigned 32-bit integers that uniquely identifies a principal relative to the IdentifierAuthority. Its length is determined by SubAuthorityCount.
}

// String returns the string representation of the RPC_SID.
func (s *RPCSID) String() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:20
	_go_fuzz_dep_.CoverTab[87397]++
											var strb strings.Builder
											strb.WriteString("S-1-")

											b := append(make([]byte, 2, 2), s.IdentifierAuthority[:]...)

											i := binary.BigEndian.Uint64(b)
											if i > math.MaxUint32 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:27
		_go_fuzz_dep_.CoverTab[87400]++
												fmt.Fprintf(&strb, "0x%s", hex.EncodeToString(s.IdentifierAuthority[:]))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:28
		// _ = "end of CoverTab[87400]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:29
		_go_fuzz_dep_.CoverTab[87401]++
												fmt.Fprintf(&strb, "%d", i)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:30
		// _ = "end of CoverTab[87401]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:31
	// _ = "end of CoverTab[87397]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:31
	_go_fuzz_dep_.CoverTab[87398]++
											for _, sub := range s.SubAuthority {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:32
		_go_fuzz_dep_.CoverTab[87402]++
												fmt.Fprintf(&strb, "-%d", sub)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:33
		// _ = "end of CoverTab[87402]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:34
	// _ = "end of CoverTab[87398]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:34
	_go_fuzz_dep_.CoverTab[87399]++
											return strb.String()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:35
	// _ = "end of CoverTab[87399]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:36
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/sid.go:36
var _ = _go_fuzz_dep_.CoverTab
