//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:1
package jwt

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:1
)

import (
	"crypto/subtle"
	"fmt"
	"time"
)

// For a type to be a Claims object, it must just have a Valid method that determines
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:9
// if the token is invalid for any supported reason
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:11
type Claims interface {
	Valid() error
}

// Structured version of Claims Section, as referenced at
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:15
// https://tools.ietf.org/html/rfc7519#section-4.1
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:15
// See examples for how to use this with your own claim types
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:18
type StandardClaims struct {
	Audience	string	`json:"aud,omitempty"`
	ExpiresAt	int64	`json:"exp,omitempty"`
	Id		string	`json:"jti,omitempty"`
	IssuedAt	int64	`json:"iat,omitempty"`
	Issuer		string	`json:"iss,omitempty"`
	NotBefore	int64	`json:"nbf,omitempty"`
	Subject		string	`json:"sub,omitempty"`
}

// Validates time based claims "exp, iat, nbf".
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:28
// There is no accounting for clock skew.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:28
// As well, if any of the above claims are not in the token, it will still
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:28
// be considered a valid claim.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:32
func (c StandardClaims) Valid() error {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:32
	_go_fuzz_dep_.CoverTab[186936]++
												vErr := new(ValidationError)
												now := TimeFunc().Unix()

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:38
	if !c.VerifyExpiresAt(now, false) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:38
		_go_fuzz_dep_.CoverTab[186941]++
													delta := time.Unix(now, 0).Sub(time.Unix(c.ExpiresAt, 0))
													vErr.Inner = fmt.Errorf("token is expired by %v", delta)
													vErr.Errors |= ValidationErrorExpired
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:41
		// _ = "end of CoverTab[186941]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:42
		_go_fuzz_dep_.CoverTab[186942]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:42
		// _ = "end of CoverTab[186942]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:42
	// _ = "end of CoverTab[186936]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:42
	_go_fuzz_dep_.CoverTab[186937]++

												if !c.VerifyIssuedAt(now, false) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:44
		_go_fuzz_dep_.CoverTab[186943]++
													vErr.Inner = fmt.Errorf("Token used before issued")
													vErr.Errors |= ValidationErrorIssuedAt
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:46
		// _ = "end of CoverTab[186943]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:47
		_go_fuzz_dep_.CoverTab[186944]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:47
		// _ = "end of CoverTab[186944]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:47
	// _ = "end of CoverTab[186937]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:47
	_go_fuzz_dep_.CoverTab[186938]++

												if !c.VerifyNotBefore(now, false) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:49
		_go_fuzz_dep_.CoverTab[186945]++
													vErr.Inner = fmt.Errorf("token is not valid yet")
													vErr.Errors |= ValidationErrorNotValidYet
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:51
		// _ = "end of CoverTab[186945]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:52
		_go_fuzz_dep_.CoverTab[186946]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:52
		// _ = "end of CoverTab[186946]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:52
	// _ = "end of CoverTab[186938]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:52
	_go_fuzz_dep_.CoverTab[186939]++

												if vErr.valid() {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:54
		_go_fuzz_dep_.CoverTab[186947]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:55
		// _ = "end of CoverTab[186947]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:56
		_go_fuzz_dep_.CoverTab[186948]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:56
		// _ = "end of CoverTab[186948]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:56
	// _ = "end of CoverTab[186939]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:56
	_go_fuzz_dep_.CoverTab[186940]++

												return vErr
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:58
	// _ = "end of CoverTab[186940]"
}

// Compares the aud claim against cmp.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:61
// If required is false, this method will return true if the value matches or is unset
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:63
func (c *StandardClaims) VerifyAudience(cmp string, req bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:63
	_go_fuzz_dep_.CoverTab[186949]++
												return verifyAud([]string{c.Audience}, cmp, req)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:64
	// _ = "end of CoverTab[186949]"
}

// Compares the exp claim against cmp.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:67
// If required is false, this method will return true if the value matches or is unset
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:69
func (c *StandardClaims) VerifyExpiresAt(cmp int64, req bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:69
	_go_fuzz_dep_.CoverTab[186950]++
												return verifyExp(c.ExpiresAt, cmp, req)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:70
	// _ = "end of CoverTab[186950]"
}

// Compares the iat claim against cmp.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:73
// If required is false, this method will return true if the value matches or is unset
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:75
func (c *StandardClaims) VerifyIssuedAt(cmp int64, req bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:75
	_go_fuzz_dep_.CoverTab[186951]++
												return verifyIat(c.IssuedAt, cmp, req)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:76
	// _ = "end of CoverTab[186951]"
}

// Compares the iss claim against cmp.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:79
// If required is false, this method will return true if the value matches or is unset
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:81
func (c *StandardClaims) VerifyIssuer(cmp string, req bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:81
	_go_fuzz_dep_.CoverTab[186952]++
												return verifyIss(c.Issuer, cmp, req)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:82
	// _ = "end of CoverTab[186952]"
}

// Compares the nbf claim against cmp.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:85
// If required is false, this method will return true if the value matches or is unset
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:87
func (c *StandardClaims) VerifyNotBefore(cmp int64, req bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:87
	_go_fuzz_dep_.CoverTab[186953]++
												return verifyNbf(c.NotBefore, cmp, req)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:88
	// _ = "end of CoverTab[186953]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:93
func verifyAud(aud []string, cmp string, required bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:93
	_go_fuzz_dep_.CoverTab[186954]++
												if len(aud) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:94
		_go_fuzz_dep_.CoverTab[186958]++
													return !required
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:95
		// _ = "end of CoverTab[186958]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:96
		_go_fuzz_dep_.CoverTab[186959]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:96
		// _ = "end of CoverTab[186959]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:96
	// _ = "end of CoverTab[186954]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:96
	_go_fuzz_dep_.CoverTab[186955]++

												result := false

												var stringClaims string
												for _, a := range aud {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:101
		_go_fuzz_dep_.CoverTab[186960]++
													if subtle.ConstantTimeCompare([]byte(a), []byte(cmp)) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:102
			_go_fuzz_dep_.CoverTab[186962]++
														result = true
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:103
			// _ = "end of CoverTab[186962]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:104
			_go_fuzz_dep_.CoverTab[186963]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:104
			// _ = "end of CoverTab[186963]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:104
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:104
		// _ = "end of CoverTab[186960]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:104
		_go_fuzz_dep_.CoverTab[186961]++
													stringClaims = stringClaims + a
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:105
		// _ = "end of CoverTab[186961]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:106
	// _ = "end of CoverTab[186955]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:106
	_go_fuzz_dep_.CoverTab[186956]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:109
	if len(stringClaims) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:109
		_go_fuzz_dep_.CoverTab[186964]++
													return !required
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:110
		// _ = "end of CoverTab[186964]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:111
		_go_fuzz_dep_.CoverTab[186965]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:111
		// _ = "end of CoverTab[186965]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:111
	// _ = "end of CoverTab[186956]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:111
	_go_fuzz_dep_.CoverTab[186957]++

												return result
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:113
	// _ = "end of CoverTab[186957]"
}

func verifyExp(exp int64, now int64, required bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:116
	_go_fuzz_dep_.CoverTab[186966]++
												if exp == 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:117
		_go_fuzz_dep_.CoverTab[186968]++
													return !required
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:118
		// _ = "end of CoverTab[186968]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:119
		_go_fuzz_dep_.CoverTab[186969]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:119
		// _ = "end of CoverTab[186969]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:119
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:119
	// _ = "end of CoverTab[186966]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:119
	_go_fuzz_dep_.CoverTab[186967]++
												return now <= exp
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:120
	// _ = "end of CoverTab[186967]"
}

func verifyIat(iat int64, now int64, required bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:123
	_go_fuzz_dep_.CoverTab[186970]++
												if iat == 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:124
		_go_fuzz_dep_.CoverTab[186972]++
													return !required
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:125
		// _ = "end of CoverTab[186972]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:126
		_go_fuzz_dep_.CoverTab[186973]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:126
		// _ = "end of CoverTab[186973]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:126
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:126
	// _ = "end of CoverTab[186970]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:126
	_go_fuzz_dep_.CoverTab[186971]++
												return now >= iat
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:127
	// _ = "end of CoverTab[186971]"
}

func verifyIss(iss string, cmp string, required bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:130
	_go_fuzz_dep_.CoverTab[186974]++
												if iss == "" {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:131
		_go_fuzz_dep_.CoverTab[186976]++
													return !required
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:132
		// _ = "end of CoverTab[186976]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:133
		_go_fuzz_dep_.CoverTab[186977]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:133
		// _ = "end of CoverTab[186977]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:133
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:133
	// _ = "end of CoverTab[186974]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:133
	_go_fuzz_dep_.CoverTab[186975]++
												if subtle.ConstantTimeCompare([]byte(iss), []byte(cmp)) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:134
		_go_fuzz_dep_.CoverTab[186978]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:135
		// _ = "end of CoverTab[186978]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:136
		_go_fuzz_dep_.CoverTab[186979]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:137
		// _ = "end of CoverTab[186979]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:138
	// _ = "end of CoverTab[186975]"
}

func verifyNbf(nbf int64, now int64, required bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:141
	_go_fuzz_dep_.CoverTab[186980]++
												if nbf == 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:142
		_go_fuzz_dep_.CoverTab[186982]++
													return !required
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:143
		// _ = "end of CoverTab[186982]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:144
		_go_fuzz_dep_.CoverTab[186983]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:144
		// _ = "end of CoverTab[186983]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:144
	// _ = "end of CoverTab[186980]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:144
	_go_fuzz_dep_.CoverTab[186981]++
												return now >= nbf
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:145
	// _ = "end of CoverTab[186981]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:146
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/claims.go:146
var _ = _go_fuzz_dep_.CoverTab
