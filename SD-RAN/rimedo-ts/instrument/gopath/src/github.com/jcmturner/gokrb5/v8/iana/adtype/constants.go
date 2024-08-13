//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/adtype/constants.go:1
// Package adtype provides Authenticator type assigned numbers.
package adtype

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/adtype/constants.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/adtype/constants.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/adtype/constants.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/adtype/constants.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/adtype/constants.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/adtype/constants.go:2
)

// Authenticator type IDs.
const (
	ADIfRelevant			int32	= 1
	ADIntendedForServer		int32	= 2
	ADIntendedForApplicationClass	int32	= 3
	ADKDCIssued			int32	= 4
	ADAndOr				int32	= 5
	ADMandatoryTicketExtensions	int32	= 6
	ADInTicketExtensions		int32	= 7
	ADMandatoryForKDC		int32	= 8
	OSFDCE				int32	= 64
	SESAME				int32	= 65
	ADOSFDCEPKICertID		int32	= 66
	ADAuthenticationStrength	int32	= 70
	ADFXFastArmor			int32	= 71
	ADFXFastUsed			int32	= 72
	ADWin2KPAC			int32	= 128
	ADEtypeNegotiation		int32	= 129
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/adtype/constants.go:23
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/adtype/constants.go:23
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/adtype/constants.go:23
var _ = _go_fuzz_dep_.CoverTab
