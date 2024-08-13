//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:1
// Package gssapi implements Generic Security Services Application Program Interface required for SPNEGO kerberos authentication.
package gssapi

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:2
)

import (
	"context"
	"fmt"

	"github.com/jcmturner/gofork/encoding/asn1"
)

// GSS-API OID names
const (
	// GSS-API OID names
	OIDKRB5		OIDName	= "KRB5"		// MechType OID for Kerberos 5
	OIDMSLegacyKRB5	OIDName	= "MSLegacyKRB5"	// MechType OID for Kerberos 5
	OIDSPNEGO	OIDName	= "SPNEGO"
	OIDGSSIAKerb	OIDName	= "GSSIAKerb"	// Indicates the client cannot get a service ticket and asks the server to serve as an intermediate to the target KDC. http://k5wiki.kerberos.org/wiki/Projects/IAKERB#IAKERB_mech
)

// GSS-API status values
const (
	StatusBadBindings	= 1 << iota
	StatusBadMech
	StatusBadName
	StatusBadNameType
	StatusBadStatus
	StatusBadSig
	StatusBadMIC
	StatusContextExpired
	StatusCredentialsExpired
	StatusDefectiveCredential
	StatusDefectiveToken
	StatusFailure
	StatusNoContext
	StatusNoCred
	StatusBadQOP
	StatusUnauthorized
	StatusUnavailable
	StatusDuplicateElement
	StatusNameNotMN
	StatusComplete
	StatusContinueNeeded
	StatusDuplicateToken
	StatusOldToken
	StatusUnseqToken
	StatusGapToken
)

// ContextToken is an interface for a GSS-API context token.
type ContextToken interface {
	Marshal() ([]byte, error)
	Unmarshal(b []byte) error
	Verify() (bool, Status)
	Context() context.Context
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:105
// Mechanism is the GSS-API interface for authentication mechanisms.
type Mechanism interface {
	OID() asn1.ObjectIdentifier
	AcquireCred() error							// acquire credentials for use (eg. AS exchange for KRB5)
	InitSecContext() (ContextToken, error)					// initiate outbound security context (eg TGS exchange builds AP_REQ to go into ContextToken to send to service)
	AcceptSecContext(ct ContextToken) (bool, context.Context, Status)	// service verifies the token server side to establish a context
	MIC() MICToken								// apply integrity check, receive as token separate from message
	VerifyMIC(mt MICToken) (bool, error)					// validate integrity check token along with message
	Wrap(msg []byte) WrapToken						// sign, optionally encrypt, encapsulate
	Unwrap(wt WrapToken) []byte						// decapsulate, decrypt if needed, validate integrity check
}

// OIDName is the type for defined GSS-API OIDs.
type OIDName string

// OID returns the OID for the provided OID name.
func (o OIDName) OID() asn1.ObjectIdentifier {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:121
	_go_fuzz_dep_.CoverTab[88831]++
												switch o {
	case OIDSPNEGO:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:123
		_go_fuzz_dep_.CoverTab[88833]++
													return asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 2}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:124
		// _ = "end of CoverTab[88833]"
	case OIDKRB5:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:125
		_go_fuzz_dep_.CoverTab[88834]++
													return asn1.ObjectIdentifier{1, 2, 840, 113554, 1, 2, 2}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:126
		// _ = "end of CoverTab[88834]"
	case OIDMSLegacyKRB5:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:127
		_go_fuzz_dep_.CoverTab[88835]++
													return asn1.ObjectIdentifier{1, 2, 840, 48018, 1, 2, 2}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:128
		// _ = "end of CoverTab[88835]"
	case OIDGSSIAKerb:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:129
		_go_fuzz_dep_.CoverTab[88836]++
													return asn1.ObjectIdentifier{1, 3, 6, 1, 5, 2, 5}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:130
		// _ = "end of CoverTab[88836]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:130
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:130
		_go_fuzz_dep_.CoverTab[88837]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:130
		// _ = "end of CoverTab[88837]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:131
	// _ = "end of CoverTab[88831]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:131
	_go_fuzz_dep_.CoverTab[88832]++
												return asn1.ObjectIdentifier{}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:132
	// _ = "end of CoverTab[88832]"
}

// Status is the GSS-API status and implements the error interface.
type Status struct {
	Code	int
	Message	string
}

// Error returns the Status description.
func (s Status) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:142
	_go_fuzz_dep_.CoverTab[88838]++
												var str string
												switch s.Code {
	case StatusBadBindings:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:145
		_go_fuzz_dep_.CoverTab[88841]++
													str = "channel binding mismatch"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:146
		// _ = "end of CoverTab[88841]"
	case StatusBadMech:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:147
		_go_fuzz_dep_.CoverTab[88842]++
													str = "unsupported mechanism requested"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:148
		// _ = "end of CoverTab[88842]"
	case StatusBadName:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:149
		_go_fuzz_dep_.CoverTab[88843]++
													str = "invalid name provided"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:150
		// _ = "end of CoverTab[88843]"
	case StatusBadNameType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:151
		_go_fuzz_dep_.CoverTab[88844]++
													str = "name of unsupported type provided"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:152
		// _ = "end of CoverTab[88844]"
	case StatusBadStatus:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:153
		_go_fuzz_dep_.CoverTab[88845]++
													str = "invalid input status selector"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:154
		// _ = "end of CoverTab[88845]"
	case StatusBadSig:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:155
		_go_fuzz_dep_.CoverTab[88846]++
													str = "token had invalid integrity check"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:156
		// _ = "end of CoverTab[88846]"
	case StatusBadMIC:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:157
		_go_fuzz_dep_.CoverTab[88847]++
													str = "preferred alias for GSS_S_BAD_SIG"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:158
		// _ = "end of CoverTab[88847]"
	case StatusContextExpired:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:159
		_go_fuzz_dep_.CoverTab[88848]++
													str = "specified security context expired"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:160
		// _ = "end of CoverTab[88848]"
	case StatusCredentialsExpired:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:161
		_go_fuzz_dep_.CoverTab[88849]++
													str = "expired credentials detected"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:162
		// _ = "end of CoverTab[88849]"
	case StatusDefectiveCredential:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:163
		_go_fuzz_dep_.CoverTab[88850]++
													str = "defective credential detected"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:164
		// _ = "end of CoverTab[88850]"
	case StatusDefectiveToken:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:165
		_go_fuzz_dep_.CoverTab[88851]++
													str = "defective token detected"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:166
		// _ = "end of CoverTab[88851]"
	case StatusFailure:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:167
		_go_fuzz_dep_.CoverTab[88852]++
													str = "failure, unspecified at GSS-API level"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:168
		// _ = "end of CoverTab[88852]"
	case StatusNoContext:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:169
		_go_fuzz_dep_.CoverTab[88853]++
													str = "no valid security context specified"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:170
		// _ = "end of CoverTab[88853]"
	case StatusNoCred:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:171
		_go_fuzz_dep_.CoverTab[88854]++
													str = "no valid credentials provided"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:172
		// _ = "end of CoverTab[88854]"
	case StatusBadQOP:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:173
		_go_fuzz_dep_.CoverTab[88855]++
													str = "unsupported QOP valu"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:174
		// _ = "end of CoverTab[88855]"
	case StatusUnauthorized:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:175
		_go_fuzz_dep_.CoverTab[88856]++
													str = "operation unauthorized"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:176
		// _ = "end of CoverTab[88856]"
	case StatusUnavailable:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:177
		_go_fuzz_dep_.CoverTab[88857]++
													str = "operation unavailable"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:178
		// _ = "end of CoverTab[88857]"
	case StatusDuplicateElement:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:179
		_go_fuzz_dep_.CoverTab[88858]++
													str = "duplicate credential element requested"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:180
		// _ = "end of CoverTab[88858]"
	case StatusNameNotMN:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:181
		_go_fuzz_dep_.CoverTab[88859]++
													str = "name contains multi-mechanism elements"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:182
		// _ = "end of CoverTab[88859]"
	case StatusComplete:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:183
		_go_fuzz_dep_.CoverTab[88860]++
													str = "normal completion"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:184
		// _ = "end of CoverTab[88860]"
	case StatusContinueNeeded:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:185
		_go_fuzz_dep_.CoverTab[88861]++
													str = "continuation call to routine required"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:186
		// _ = "end of CoverTab[88861]"
	case StatusDuplicateToken:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:187
		_go_fuzz_dep_.CoverTab[88862]++
													str = "duplicate per-message token detected"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:188
		// _ = "end of CoverTab[88862]"
	case StatusOldToken:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:189
		_go_fuzz_dep_.CoverTab[88863]++
													str = "timed-out per-message token detected"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:190
		// _ = "end of CoverTab[88863]"
	case StatusUnseqToken:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:191
		_go_fuzz_dep_.CoverTab[88864]++
													str = "reordered (early) per-message token detected"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:192
		// _ = "end of CoverTab[88864]"
	case StatusGapToken:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:193
		_go_fuzz_dep_.CoverTab[88865]++
													str = "skipped predecessor token(s) detected"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:194
		// _ = "end of CoverTab[88865]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:195
		_go_fuzz_dep_.CoverTab[88866]++
													str = "unknown GSS-API error status"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:196
		// _ = "end of CoverTab[88866]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:197
	// _ = "end of CoverTab[88838]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:197
	_go_fuzz_dep_.CoverTab[88839]++
												if s.Message != "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:198
		_go_fuzz_dep_.CoverTab[88867]++
													return fmt.Sprintf("%s: %s", str, s.Message)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:199
		// _ = "end of CoverTab[88867]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:200
		_go_fuzz_dep_.CoverTab[88868]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:200
		// _ = "end of CoverTab[88868]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:200
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:200
	// _ = "end of CoverTab[88839]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:200
	_go_fuzz_dep_.CoverTab[88840]++
												return str
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:201
	// _ = "end of CoverTab[88840]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:202
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/gssapi/gssapi.go:202
var _ = _go_fuzz_dep_.CoverTab
