//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/nametype/constants.go:1
// Package nametype provides Kerberos 5 principal name type numbers.
package nametype

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/nametype/constants.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/nametype/constants.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/nametype/constants.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/nametype/constants.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/nametype/constants.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/nametype/constants.go:2
)

// Kerberos name type IDs.
const (
	KRB_NT_UNKNOWN		int32	= 0	//Name type not known
	KRB_NT_PRINCIPAL	int32	= 1	//Just the name of the principal as in DCE,  or for users
	KRB_NT_SRV_INST		int32	= 2	//Service and other unique instance (krbtgt)
	KRB_NT_SRV_HST		int32	= 3	//Service with host name as instance (telnet, rcommands)
	KRB_NT_SRV_XHST		int32	= 4	//Service with host as remaining components
	KRB_NT_UID		int32	= 5	//Unique ID
	KRB_NT_X500_PRINCIPAL	int32	= 6	//Encoded X.509 Distinguished name [RFC2253]
	KRB_NT_SMTP_NAME	int32	= 7	//Name in form of SMTP email name (e.g., user@example.com)
	KRB_NT_ENTERPRISE	int32	= 10	//Enterprise name; may be mapped to principal name
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/nametype/constants.go:15
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/nametype/constants.go:15
var _ = _go_fuzz_dep_.CoverTab
