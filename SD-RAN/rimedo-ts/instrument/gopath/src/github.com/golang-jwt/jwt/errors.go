//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:1
package jwt

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:1
)

import (
	"errors"
)

// Error constants
var (
	ErrInvalidKey		= errors.New("key is invalid")
	ErrInvalidKeyType	= errors.New("key is of invalid type")
	ErrHashUnavailable	= errors.New("the requested hash function is unavailable")
)

// The errors that might occur when parsing and validating a token
const (
	ValidationErrorMalformed	uint32	= 1 << iota	// Token is malformed
	ValidationErrorUnverifiable				// Token could not be verified because of signing problems
	ValidationErrorSignatureInvalid				// Signature validation failed

	// Standard Claim validation errors
	ValidationErrorAudience		// AUD validation failed
	ValidationErrorExpired		// EXP validation failed
	ValidationErrorIssuedAt		// IAT validation failed
	ValidationErrorIssuer		// ISS validation failed
	ValidationErrorNotValidYet	// NBF validation failed
	ValidationErrorId		// JTI validation failed
	ValidationErrorClaimsInvalid	// Generic claims validation error
)

// Helper for constructing a ValidationError with a string error message
func NewValidationError(errorText string, errorFlags uint32) *ValidationError {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:31
	_go_fuzz_dep_.CoverTab[187081]++
												return &ValidationError{
		text:	errorText,
		Errors:	errorFlags,
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:35
	// _ = "end of CoverTab[187081]"
}

// The error from Parse if token is not valid
type ValidationError struct {
	Inner	error	// stores the error returned by external dependencies, i.e.: KeyFunc
	Errors	uint32	// bitfield.  see ValidationError... constants
	text	string	// errors that do not have a valid error just have text
}

// Validation error is an error type
func (e ValidationError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:46
	_go_fuzz_dep_.CoverTab[187082]++
												if e.Inner != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:47
		_go_fuzz_dep_.CoverTab[187083]++
													return e.Inner.Error()
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:48
		// _ = "end of CoverTab[187083]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:49
		_go_fuzz_dep_.CoverTab[187084]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:49
		if e.text != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:49
			_go_fuzz_dep_.CoverTab[187085]++
														return e.text
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:50
			// _ = "end of CoverTab[187085]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:51
			_go_fuzz_dep_.CoverTab[187086]++
														return "token is invalid"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:52
			// _ = "end of CoverTab[187086]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:53
		// _ = "end of CoverTab[187084]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:53
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:53
	// _ = "end of CoverTab[187082]"
}

// No errors
func (e *ValidationError) valid() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:57
	_go_fuzz_dep_.CoverTab[187087]++
												return e.Errors == 0
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:58
	// _ = "end of CoverTab[187087]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:59
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/errors.go:59
var _ = _go_fuzz_dep_.CoverTab
