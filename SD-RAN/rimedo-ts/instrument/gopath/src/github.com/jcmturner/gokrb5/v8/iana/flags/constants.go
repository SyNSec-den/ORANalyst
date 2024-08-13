//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/flags/constants.go:1
// Package flags provides Kerberos 5 flag assigned numbers.
package flags

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/flags/constants.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/flags/constants.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/flags/constants.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/flags/constants.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/flags/constants.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/flags/constants.go:2
)

// Flag values for KRB5 messages and tickets.
const (
	Reserved		= 0
	Forwardable		= 1
	Forwarded		= 2
	Proxiable		= 3
	Proxy			= 4
	AllowPostDate		= 5
	MayPostDate		= 5
	PostDated		= 6
	Invalid			= 7
	Renewable		= 8
	Initial			= 9
	PreAuthent		= 10
	HWAuthent		= 11
	OptHardwareAuth		= 11
	RequestAnonymous	= 12
	TransitedPolicyChecked	= 12
	OKAsDelegate		= 13
	EncPARep		= 15
	Canonicalize		= 15
	DisableTransitedCheck	= 26
	RenewableOK		= 27
	EncTktInSkey		= 28
	Renew			= 30
	Validate		= 31

	// AP Option Flags
	// 0 Reserved for future use.
	APOptionUseSessionKey	= 1
	APOptionMutualRequired	= 2
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/flags/constants.go:36
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/flags/constants.go:36
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/flags/constants.go:36
var _ = _go_fuzz_dep_.CoverTab
