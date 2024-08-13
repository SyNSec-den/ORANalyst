//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:1
package kadmin

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:1
)

import (
	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/types"
)

// ChangePasswdData is the payload to a password change message.
type ChangePasswdData struct {
	NewPasswd	[]byte			`asn1:"explicit,tag:0"`
	TargName	types.PrincipalName	`asn1:"explicit,optional,tag:1"`
	TargRealm	string			`asn1:"generalstring,optional,explicit,tag:2"`
}

// Marshal ChangePasswdData into a byte slice.
func (c *ChangePasswdData) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:16
	_go_fuzz_dep_.CoverTab[88180]++
													b, err := asn1.Marshal(*c)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:18
		_go_fuzz_dep_.CoverTab[88182]++
														return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:19
		// _ = "end of CoverTab[88182]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:20
		_go_fuzz_dep_.CoverTab[88183]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:20
		// _ = "end of CoverTab[88183]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:20
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:20
	// _ = "end of CoverTab[88180]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:20
	_go_fuzz_dep_.CoverTab[88181]++

													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:22
	// _ = "end of CoverTab[88181]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:23
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/kadmin/changepasswddata.go:23
var _ = _go_fuzz_dep_.CoverTab
