//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:1
package pac

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:1
)

import (
	"bytes"
	"fmt"

	"github.com/jcmturner/rpc/v2/mstypes"
	"github.com/jcmturner/rpc/v2/ndr"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:13
// DeviceClaimsInfo implements https://msdn.microsoft.com/en-us/library/hh554226.aspx
type DeviceClaimsInfo struct {
	ClaimsSetMetadata	mstypes.ClaimsSetMetadata
	ClaimsSet		mstypes.ClaimsSet
}

// Unmarshal bytes into the ClientClaimsInfo struct
func (k *DeviceClaimsInfo) Unmarshal(b []byte) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:20
	_go_fuzz_dep_.CoverTab[87447]++
													dec := ndr.NewDecoder(bytes.NewReader(b))
													m := new(mstypes.ClaimsSetMetadata)
													err = dec.Decode(m)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:24
		_go_fuzz_dep_.CoverTab[87450]++
														err = fmt.Errorf("error unmarshaling ClientClaimsInfo ClaimsSetMetadata: %v", err)
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:26
		// _ = "end of CoverTab[87450]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:27
		_go_fuzz_dep_.CoverTab[87451]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:27
		// _ = "end of CoverTab[87451]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:27
	// _ = "end of CoverTab[87447]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:27
	_go_fuzz_dep_.CoverTab[87448]++
													k.ClaimsSetMetadata = *m
													k.ClaimsSet, err = k.ClaimsSetMetadata.ClaimsSet()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:30
		_go_fuzz_dep_.CoverTab[87452]++
														err = fmt.Errorf("error unmarshaling ClientClaimsInfo ClaimsSet: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:31
		// _ = "end of CoverTab[87452]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:32
		_go_fuzz_dep_.CoverTab[87453]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:32
		// _ = "end of CoverTab[87453]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:32
	// _ = "end of CoverTab[87448]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:32
	_go_fuzz_dep_.CoverTab[87449]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:33
	// _ = "end of CoverTab[87449]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:34
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/pac/device_claims.go:34
var _ = _go_fuzz_dep_.CoverTab
