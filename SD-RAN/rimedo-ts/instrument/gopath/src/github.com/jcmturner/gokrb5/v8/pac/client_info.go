//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:1
package pac

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:1
)

import (
	"bytes"

	"github.com/jcmturner/rpc/v2/mstypes"
)

// ClientInfo implements https://msdn.microsoft.com/en-us/library/cc237951.aspx
type ClientInfo struct {
	ClientID	mstypes.FileTime	// A FILETIME structure in little-endian format that contains the Kerberos initial ticket-granting ticket TGT authentication time
	NameLength	uint16			// An unsigned 16-bit integer in little-endian format that specifies the length, in bytes, of the Name field.
	Name		string			// An array of 16-bit Unicode characters in little-endian format that contains the client's account name.
}

// Unmarshal bytes into the ClientInfo struct
func (k *ClientInfo) Unmarshal(b []byte) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:17
	_go_fuzz_dep_.CoverTab[87410]++

												r := mstypes.NewReader(bytes.NewReader(b))

												k.ClientID, err = r.FileTime()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:22
		_go_fuzz_dep_.CoverTab[87413]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:23
		// _ = "end of CoverTab[87413]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:24
		_go_fuzz_dep_.CoverTab[87414]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:24
		// _ = "end of CoverTab[87414]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:24
	// _ = "end of CoverTab[87410]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:24
	_go_fuzz_dep_.CoverTab[87411]++
												k.NameLength, err = r.Uint16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:26
		_go_fuzz_dep_.CoverTab[87415]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:27
		// _ = "end of CoverTab[87415]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:28
		_go_fuzz_dep_.CoverTab[87416]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:28
		// _ = "end of CoverTab[87416]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:28
	// _ = "end of CoverTab[87411]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:28
	_go_fuzz_dep_.CoverTab[87412]++
												k.Name, err = r.UTF16String(int(k.NameLength))
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:30
	// _ = "end of CoverTab[87412]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:31
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_info.go:31
var _ = _go_fuzz_dep_.CoverTab
