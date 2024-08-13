//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:1
package jwt

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:1
)

import (
	"encoding/json"
	"errors"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:7
)

// Claims type that uses the map[string]interface{} for JSON decoding
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:9
// This is the default claims type if you don't supply one
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:11
type MapClaims map[string]interface{}

// VerifyAudience Compares the aud claim against cmp.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:13
// If required is false, this method will return true if the value matches or is unset
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:15
func (m MapClaims) VerifyAudience(cmp string, req bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:15
	_go_fuzz_dep_.CoverTab[187109]++
													var aud []string
													switch v := m["aud"].(type) {
	case string:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:18
		_go_fuzz_dep_.CoverTab[187111]++
														aud = append(aud, v)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:19
		// _ = "end of CoverTab[187111]"
	case []string:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:20
		_go_fuzz_dep_.CoverTab[187112]++
														aud = v
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:21
		// _ = "end of CoverTab[187112]"
	case []interface{}:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:22
		_go_fuzz_dep_.CoverTab[187113]++
														for _, a := range v {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:23
			_go_fuzz_dep_.CoverTab[187114]++
															vs, ok := a.(string)
															if !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:25
				_go_fuzz_dep_.CoverTab[187116]++
																return false
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:26
				// _ = "end of CoverTab[187116]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:27
				_go_fuzz_dep_.CoverTab[187117]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:27
				// _ = "end of CoverTab[187117]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:27
			}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:27
			// _ = "end of CoverTab[187114]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:27
			_go_fuzz_dep_.CoverTab[187115]++
															aud = append(aud, vs)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:28
			// _ = "end of CoverTab[187115]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:29
		// _ = "end of CoverTab[187113]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:30
	// _ = "end of CoverTab[187109]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:30
	_go_fuzz_dep_.CoverTab[187110]++
													return verifyAud(aud, cmp, req)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:31
	// _ = "end of CoverTab[187110]"
}

// Compares the exp claim against cmp.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:34
// If required is false, this method will return true if the value matches or is unset
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:36
func (m MapClaims) VerifyExpiresAt(cmp int64, req bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:36
	_go_fuzz_dep_.CoverTab[187118]++
													exp, ok := m["exp"]
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:38
		_go_fuzz_dep_.CoverTab[187121]++
														return !req
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:39
		// _ = "end of CoverTab[187121]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:40
		_go_fuzz_dep_.CoverTab[187122]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:40
		// _ = "end of CoverTab[187122]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:40
	// _ = "end of CoverTab[187118]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:40
	_go_fuzz_dep_.CoverTab[187119]++
													switch expType := exp.(type) {
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:42
		_go_fuzz_dep_.CoverTab[187123]++
														return verifyExp(int64(expType), cmp, req)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:43
		// _ = "end of CoverTab[187123]"
	case json.Number:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:44
		_go_fuzz_dep_.CoverTab[187124]++
														v, _ := expType.Int64()
														return verifyExp(v, cmp, req)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:46
		// _ = "end of CoverTab[187124]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:47
	// _ = "end of CoverTab[187119]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:47
	_go_fuzz_dep_.CoverTab[187120]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:48
	// _ = "end of CoverTab[187120]"
}

// Compares the iat claim against cmp.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:51
// If required is false, this method will return true if the value matches or is unset
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:53
func (m MapClaims) VerifyIssuedAt(cmp int64, req bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:53
	_go_fuzz_dep_.CoverTab[187125]++
													iat, ok := m["iat"]
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:55
		_go_fuzz_dep_.CoverTab[187128]++
														return !req
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:56
		// _ = "end of CoverTab[187128]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:57
		_go_fuzz_dep_.CoverTab[187129]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:57
		// _ = "end of CoverTab[187129]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:57
	// _ = "end of CoverTab[187125]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:57
	_go_fuzz_dep_.CoverTab[187126]++
													switch iatType := iat.(type) {
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:59
		_go_fuzz_dep_.CoverTab[187130]++
														return verifyIat(int64(iatType), cmp, req)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:60
		// _ = "end of CoverTab[187130]"
	case json.Number:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:61
		_go_fuzz_dep_.CoverTab[187131]++
														v, _ := iatType.Int64()
														return verifyIat(v, cmp, req)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:63
		// _ = "end of CoverTab[187131]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:64
	// _ = "end of CoverTab[187126]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:64
	_go_fuzz_dep_.CoverTab[187127]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:65
	// _ = "end of CoverTab[187127]"
}

// Compares the iss claim against cmp.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:68
// If required is false, this method will return true if the value matches or is unset
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:70
func (m MapClaims) VerifyIssuer(cmp string, req bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:70
	_go_fuzz_dep_.CoverTab[187132]++
													iss, _ := m["iss"].(string)
													return verifyIss(iss, cmp, req)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:72
	// _ = "end of CoverTab[187132]"
}

// Compares the nbf claim against cmp.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:75
// If required is false, this method will return true if the value matches or is unset
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:77
func (m MapClaims) VerifyNotBefore(cmp int64, req bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:77
	_go_fuzz_dep_.CoverTab[187133]++
													nbf, ok := m["nbf"]
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:79
		_go_fuzz_dep_.CoverTab[187136]++
														return !req
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:80
		// _ = "end of CoverTab[187136]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:81
		_go_fuzz_dep_.CoverTab[187137]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:81
		// _ = "end of CoverTab[187137]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:81
	// _ = "end of CoverTab[187133]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:81
	_go_fuzz_dep_.CoverTab[187134]++
													switch nbfType := nbf.(type) {
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:83
		_go_fuzz_dep_.CoverTab[187138]++
														return verifyNbf(int64(nbfType), cmp, req)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:84
		// _ = "end of CoverTab[187138]"
	case json.Number:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:85
		_go_fuzz_dep_.CoverTab[187139]++
														v, _ := nbfType.Int64()
														return verifyNbf(v, cmp, req)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:87
		// _ = "end of CoverTab[187139]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:88
	// _ = "end of CoverTab[187134]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:88
	_go_fuzz_dep_.CoverTab[187135]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:89
	// _ = "end of CoverTab[187135]"
}

// Validates time based claims "exp, iat, nbf".
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:92
// There is no accounting for clock skew.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:92
// As well, if any of the above claims are not in the token, it will still
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:92
// be considered a valid claim.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:96
func (m MapClaims) Valid() error {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:96
	_go_fuzz_dep_.CoverTab[187140]++
													vErr := new(ValidationError)
													now := TimeFunc().Unix()

													if !m.VerifyExpiresAt(now, false) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:100
		_go_fuzz_dep_.CoverTab[187145]++
														vErr.Inner = errors.New("Token is expired")
														vErr.Errors |= ValidationErrorExpired
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:102
		// _ = "end of CoverTab[187145]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:103
		_go_fuzz_dep_.CoverTab[187146]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:103
		// _ = "end of CoverTab[187146]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:103
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:103
	// _ = "end of CoverTab[187140]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:103
	_go_fuzz_dep_.CoverTab[187141]++

													if !m.VerifyIssuedAt(now, false) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:105
		_go_fuzz_dep_.CoverTab[187147]++
														vErr.Inner = errors.New("Token used before issued")
														vErr.Errors |= ValidationErrorIssuedAt
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:107
		// _ = "end of CoverTab[187147]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:108
		_go_fuzz_dep_.CoverTab[187148]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:108
		// _ = "end of CoverTab[187148]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:108
	// _ = "end of CoverTab[187141]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:108
	_go_fuzz_dep_.CoverTab[187142]++

													if !m.VerifyNotBefore(now, false) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:110
		_go_fuzz_dep_.CoverTab[187149]++
														vErr.Inner = errors.New("Token is not valid yet")
														vErr.Errors |= ValidationErrorNotValidYet
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:112
		// _ = "end of CoverTab[187149]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:113
		_go_fuzz_dep_.CoverTab[187150]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:113
		// _ = "end of CoverTab[187150]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:113
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:113
	// _ = "end of CoverTab[187142]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:113
	_go_fuzz_dep_.CoverTab[187143]++

													if vErr.valid() {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:115
		_go_fuzz_dep_.CoverTab[187151]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:116
		// _ = "end of CoverTab[187151]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:117
		_go_fuzz_dep_.CoverTab[187152]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:117
		// _ = "end of CoverTab[187152]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:117
	// _ = "end of CoverTab[187143]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:117
	_go_fuzz_dep_.CoverTab[187144]++

													return vErr
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:119
	// _ = "end of CoverTab[187144]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:120
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/map_claims.go:120
var _ = _go_fuzz_dep_.CoverTab
