//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:1
package pac

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:1
)

import (
	"bytes"
	"fmt"

	"github.com/jcmturner/rpc/v2/mstypes"
	"github.com/jcmturner/rpc/v2/ndr"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:13
// ClientClaimsInfo implements https://msdn.microsoft.com/en-us/library/hh536365.aspx
type ClientClaimsInfo struct {
	ClaimsSetMetadata	mstypes.ClaimsSetMetadata
	ClaimsSet		mstypes.ClaimsSet
}

// Unmarshal bytes into the ClientClaimsInfo struct
func (k *ClientClaimsInfo) Unmarshal(b []byte) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:20
	_go_fuzz_dep_.CoverTab[87403]++
													dec := ndr.NewDecoder(bytes.NewReader(b))
													m := new(mstypes.ClaimsSetMetadata)
													err = dec.Decode(m)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:24
		_go_fuzz_dep_.CoverTab[87406]++
														err = fmt.Errorf("error unmarshaling ClientClaimsInfo ClaimsSetMetadata: %v", err)
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:26
		// _ = "end of CoverTab[87406]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:27
		_go_fuzz_dep_.CoverTab[87407]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:27
		// _ = "end of CoverTab[87407]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:27
	// _ = "end of CoverTab[87403]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:27
	_go_fuzz_dep_.CoverTab[87404]++
													k.ClaimsSetMetadata = *m
													k.ClaimsSet, err = k.ClaimsSetMetadata.ClaimsSet()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:30
		_go_fuzz_dep_.CoverTab[87408]++
														err = fmt.Errorf("error unmarshaling ClientClaimsInfo ClaimsSet: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:31
		// _ = "end of CoverTab[87408]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:32
		_go_fuzz_dep_.CoverTab[87409]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:32
		// _ = "end of CoverTab[87409]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:32
	// _ = "end of CoverTab[87404]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:32
	_go_fuzz_dep_.CoverTab[87405]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:33
	// _ = "end of CoverTab[87405]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:34
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/client_claims.go:34
var _ = _go_fuzz_dep_.CoverTab
