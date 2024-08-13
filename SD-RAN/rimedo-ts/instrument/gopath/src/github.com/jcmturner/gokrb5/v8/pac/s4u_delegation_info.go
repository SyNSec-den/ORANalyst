//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:1
package pac

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:1
)

import (
	"bytes"
	"fmt"

	"github.com/jcmturner/rpc/v2/mstypes"
	"github.com/jcmturner/rpc/v2/ndr"
)

// S4UDelegationInfo implements https://msdn.microsoft.com/en-us/library/cc237944.aspx
type S4UDelegationInfo struct {
	S4U2proxyTarget		mstypes.RPCUnicodeString	// The name of the principal to whom the application can forward the ticket.
	TransitedListSize	uint32
	S4UTransitedServices	[]mstypes.RPCUnicodeString	`ndr:"pointer,conformant"`	// List of all services that have been delegated through by this client and subsequent services or servers.. Size is value of TransitedListSize
}

// Unmarshal bytes into the S4UDelegationInfo struct
func (k *S4UDelegationInfo) Unmarshal(b []byte) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:19
	_go_fuzz_dep_.CoverTab[87589]++
													dec := ndr.NewDecoder(bytes.NewReader(b))
													err = dec.Decode(k)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:22
		_go_fuzz_dep_.CoverTab[87591]++
														err = fmt.Errorf("error unmarshaling S4UDelegationInfo: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:23
		// _ = "end of CoverTab[87591]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:24
		_go_fuzz_dep_.CoverTab[87592]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:24
		// _ = "end of CoverTab[87592]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:24
	// _ = "end of CoverTab[87589]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:24
	_go_fuzz_dep_.CoverTab[87590]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:25
	// _ = "end of CoverTab[87590]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:26
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/s4u_delegation_info.go:26
var _ = _go_fuzz_dep_.CoverTab
