//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/keyusage/constants.go:1
// Package keyusage provides Kerberos 5 key usage assigned numbers.
package keyusage

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/keyusage/constants.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/keyusage/constants.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/keyusage/constants.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/keyusage/constants.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/keyusage/constants.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/keyusage/constants.go:2
)

// Key usage numbers.
const (
	AS_REQ_PA_ENC_TIMESTAMP				= 1
	KDC_REP_TICKET					= 2
	AS_REP_ENCPART					= 3
	TGS_REQ_KDC_REQ_BODY_AUTHDATA_SESSION_KEY	= 4
	TGS_REQ_KDC_REQ_BODY_AUTHDATA_SUB_KEY		= 5
	TGS_REQ_PA_TGS_REQ_AP_REQ_AUTHENTICATOR_CHKSUM	= 6
	TGS_REQ_PA_TGS_REQ_AP_REQ_AUTHENTICATOR		= 7
	TGS_REP_ENCPART_SESSION_KEY			= 8
	TGS_REP_ENCPART_AUTHENTICATOR_SUB_KEY		= 9
	AP_REQ_AUTHENTICATOR_CHKSUM			= 10
	AP_REQ_AUTHENTICATOR				= 11
	AP_REP_ENCPART					= 12
	KRB_PRIV_ENCPART				= 13
	KRB_CRED_ENCPART				= 14
	KRB_SAFE_CHKSUM					= 15
	KERB_NON_KERB_SALT				= 16
	KERB_NON_KERB_CKSUM_SALT			= 17
	//18.  Reserved for future use in Kerberos and related protocols.
	AD_KDC_ISSUED_CHKSUM	= 19
	//20-21.  Reserved for future use in Kerberos and related protocols.
	GSSAPI_ACCEPTOR_SEAL		= 22
	GSSAPI_ACCEPTOR_SIGN		= 23
	GSSAPI_INITIATOR_SEAL		= 24
	GSSAPI_INITIATOR_SIGN		= 25
	KEY_USAGE_FAST_REQ_CHKSUM	= 50
	KEY_USAGE_FAST_ENC		= 51
	KEY_USAGE_FAST_REP		= 52
	KEY_USAGE_FAST_FINISHED		= 53
	KEY_USAGE_ENC_CHALLENGE_CLIENT	= 54
	KEY_USAGE_ENC_CHALLENGE_KDC	= 55
	KEY_USAGE_AS_REQ		= 56
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/keyusage/constants.go:42
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/keyusage/constants.go:42
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/iana/keyusage/constants.go:42
var _ = _go_fuzz_dep_.CoverTab
