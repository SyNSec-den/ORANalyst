//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:1
package jwt

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:1
)

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Parser struct {
	ValidMethods		[]string	// If populated, only these methods will be considered valid
	UseJSONNumber		bool		// Use JSON Number format in JSON decoder
	SkipClaimsValidation	bool		// Skip claims validation during token parsing
}

// Parse, validate, and return a token.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:16
// keyFunc will receive the parsed token and should return the key for validating.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:16
// If everything is kosher, err will be nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:19
func (p *Parser) Parse(tokenString string, keyFunc Keyfunc) (*Token, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:19
	_go_fuzz_dep_.CoverTab[187165]++
												return p.ParseWithClaims(tokenString, MapClaims{}, keyFunc)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:20
	// _ = "end of CoverTab[187165]"
}

func (p *Parser) ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:23
	_go_fuzz_dep_.CoverTab[187166]++
												token, parts, err := p.ParseUnverified(tokenString, claims)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:25
		_go_fuzz_dep_.CoverTab[187174]++
													return token, err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:26
		// _ = "end of CoverTab[187174]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:27
		_go_fuzz_dep_.CoverTab[187175]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:27
		// _ = "end of CoverTab[187175]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:27
	// _ = "end of CoverTab[187166]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:27
	_go_fuzz_dep_.CoverTab[187167]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:30
	if p.ValidMethods != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:30
		_go_fuzz_dep_.CoverTab[187176]++
													var signingMethodValid = false
													var alg = token.Method.Alg()
													for _, m := range p.ValidMethods {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:33
			_go_fuzz_dep_.CoverTab[187178]++
														if m == alg {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:34
				_go_fuzz_dep_.CoverTab[187179]++
															signingMethodValid = true
															break
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:36
				// _ = "end of CoverTab[187179]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:37
				_go_fuzz_dep_.CoverTab[187180]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:37
				// _ = "end of CoverTab[187180]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:37
			}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:37
			// _ = "end of CoverTab[187178]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:38
		// _ = "end of CoverTab[187176]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:38
		_go_fuzz_dep_.CoverTab[187177]++
													if !signingMethodValid {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:39
			_go_fuzz_dep_.CoverTab[187181]++

														return token, NewValidationError(fmt.Sprintf("signing method %v is invalid", alg), ValidationErrorSignatureInvalid)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:41
			// _ = "end of CoverTab[187181]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:42
			_go_fuzz_dep_.CoverTab[187182]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:42
			// _ = "end of CoverTab[187182]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:42
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:42
		// _ = "end of CoverTab[187177]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:43
		_go_fuzz_dep_.CoverTab[187183]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:43
		// _ = "end of CoverTab[187183]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:43
	// _ = "end of CoverTab[187167]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:43
	_go_fuzz_dep_.CoverTab[187168]++

	// Lookup key
	var key interface{}
	if keyFunc == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:47
		_go_fuzz_dep_.CoverTab[187184]++

													return token, NewValidationError("no Keyfunc was provided.", ValidationErrorUnverifiable)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:49
		// _ = "end of CoverTab[187184]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:50
		_go_fuzz_dep_.CoverTab[187185]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:50
		// _ = "end of CoverTab[187185]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:50
	// _ = "end of CoverTab[187168]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:50
	_go_fuzz_dep_.CoverTab[187169]++
												if key, err = keyFunc(token); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:51
		_go_fuzz_dep_.CoverTab[187186]++

													if ve, ok := err.(*ValidationError); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:53
			_go_fuzz_dep_.CoverTab[187188]++
														return token, ve
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:54
			// _ = "end of CoverTab[187188]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:55
			_go_fuzz_dep_.CoverTab[187189]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:55
			// _ = "end of CoverTab[187189]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:55
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:55
		// _ = "end of CoverTab[187186]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:55
		_go_fuzz_dep_.CoverTab[187187]++
													return token, &ValidationError{Inner: err, Errors: ValidationErrorUnverifiable}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:56
		// _ = "end of CoverTab[187187]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:57
		_go_fuzz_dep_.CoverTab[187190]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:57
		// _ = "end of CoverTab[187190]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:57
	// _ = "end of CoverTab[187169]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:57
	_go_fuzz_dep_.CoverTab[187170]++

												vErr := &ValidationError{}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:62
	if !p.SkipClaimsValidation {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:62
		_go_fuzz_dep_.CoverTab[187191]++
													if err := token.Claims.Valid(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:63
			_go_fuzz_dep_.CoverTab[187192]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:67
			if e, ok := err.(*ValidationError); !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:67
				_go_fuzz_dep_.CoverTab[187193]++
															vErr = &ValidationError{Inner: err, Errors: ValidationErrorClaimsInvalid}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:68
				// _ = "end of CoverTab[187193]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:69
				_go_fuzz_dep_.CoverTab[187194]++
															vErr = e
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:70
				// _ = "end of CoverTab[187194]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:71
			// _ = "end of CoverTab[187192]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:72
			_go_fuzz_dep_.CoverTab[187195]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:72
			// _ = "end of CoverTab[187195]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:72
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:72
		// _ = "end of CoverTab[187191]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:73
		_go_fuzz_dep_.CoverTab[187196]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:73
		// _ = "end of CoverTab[187196]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:73
	// _ = "end of CoverTab[187170]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:73
	_go_fuzz_dep_.CoverTab[187171]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:76
	token.Signature = parts[2]
	if err = token.Method.Verify(strings.Join(parts[0:2], "."), token.Signature, key); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:77
		_go_fuzz_dep_.CoverTab[187197]++
													vErr.Inner = err
													vErr.Errors |= ValidationErrorSignatureInvalid
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:79
		// _ = "end of CoverTab[187197]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:80
		_go_fuzz_dep_.CoverTab[187198]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:80
		// _ = "end of CoverTab[187198]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:80
	// _ = "end of CoverTab[187171]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:80
	_go_fuzz_dep_.CoverTab[187172]++

												if vErr.valid() {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:82
		_go_fuzz_dep_.CoverTab[187199]++
													token.Valid = true
													return token, nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:84
		// _ = "end of CoverTab[187199]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:85
		_go_fuzz_dep_.CoverTab[187200]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:85
		// _ = "end of CoverTab[187200]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:85
	// _ = "end of CoverTab[187172]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:85
	_go_fuzz_dep_.CoverTab[187173]++

												return token, vErr
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:87
	// _ = "end of CoverTab[187173]"
}

// WARNING: Don't use this method unless you know what you're doing
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:90
//
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:90
// This method parses the token but doesn't validate the signature. It's only
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:90
// ever useful in cases where you know the signature is valid (because it has
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:90
// been checked previously in the stack) and you want to extract values from
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:90
// it.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:96
func (p *Parser) ParseUnverified(tokenString string, claims Claims) (token *Token, parts []string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:96
	_go_fuzz_dep_.CoverTab[187201]++
												parts = strings.Split(tokenString, ".")
												if len(parts) != 3 {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:98
		_go_fuzz_dep_.CoverTab[187210]++
													return nil, parts, NewValidationError("token contains an invalid number of segments", ValidationErrorMalformed)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:99
		// _ = "end of CoverTab[187210]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:100
		_go_fuzz_dep_.CoverTab[187211]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:100
		// _ = "end of CoverTab[187211]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:100
	// _ = "end of CoverTab[187201]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:100
	_go_fuzz_dep_.CoverTab[187202]++

												token = &Token{Raw: tokenString}

	// parse Header
	var headerBytes []byte
	if headerBytes, err = DecodeSegment(parts[0]); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:106
		_go_fuzz_dep_.CoverTab[187212]++
													if strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:107
			_go_fuzz_dep_.CoverTab[187214]++
														return token, parts, NewValidationError("tokenstring should not contain 'bearer '", ValidationErrorMalformed)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:108
			// _ = "end of CoverTab[187214]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:109
			_go_fuzz_dep_.CoverTab[187215]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:109
			// _ = "end of CoverTab[187215]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:109
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:109
		// _ = "end of CoverTab[187212]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:109
		_go_fuzz_dep_.CoverTab[187213]++
													return token, parts, &ValidationError{Inner: err, Errors: ValidationErrorMalformed}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:110
		// _ = "end of CoverTab[187213]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:111
		_go_fuzz_dep_.CoverTab[187216]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:111
		// _ = "end of CoverTab[187216]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:111
	// _ = "end of CoverTab[187202]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:111
	_go_fuzz_dep_.CoverTab[187203]++
												if err = json.Unmarshal(headerBytes, &token.Header); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:112
		_go_fuzz_dep_.CoverTab[187217]++
													return token, parts, &ValidationError{Inner: err, Errors: ValidationErrorMalformed}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:113
		// _ = "end of CoverTab[187217]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:114
		_go_fuzz_dep_.CoverTab[187218]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:114
		// _ = "end of CoverTab[187218]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:114
	// _ = "end of CoverTab[187203]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:114
	_go_fuzz_dep_.CoverTab[187204]++

	// parse Claims
	var claimBytes []byte
	token.Claims = claims

	if claimBytes, err = DecodeSegment(parts[1]); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:120
		_go_fuzz_dep_.CoverTab[187219]++
													return token, parts, &ValidationError{Inner: err, Errors: ValidationErrorMalformed}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:121
		// _ = "end of CoverTab[187219]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:122
		_go_fuzz_dep_.CoverTab[187220]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:122
		// _ = "end of CoverTab[187220]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:122
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:122
	// _ = "end of CoverTab[187204]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:122
	_go_fuzz_dep_.CoverTab[187205]++
												dec := json.NewDecoder(bytes.NewBuffer(claimBytes))
												if p.UseJSONNumber {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:124
		_go_fuzz_dep_.CoverTab[187221]++
													dec.UseNumber()
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:125
		// _ = "end of CoverTab[187221]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:126
		_go_fuzz_dep_.CoverTab[187222]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:126
		// _ = "end of CoverTab[187222]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:126
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:126
	// _ = "end of CoverTab[187205]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:126
	_go_fuzz_dep_.CoverTab[187206]++

												if c, ok := token.Claims.(MapClaims); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:128
		_go_fuzz_dep_.CoverTab[187223]++
													err = dec.Decode(&c)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:129
		// _ = "end of CoverTab[187223]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:130
		_go_fuzz_dep_.CoverTab[187224]++
													err = dec.Decode(&claims)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:131
		// _ = "end of CoverTab[187224]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:132
	// _ = "end of CoverTab[187206]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:132
	_go_fuzz_dep_.CoverTab[187207]++

												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:134
		_go_fuzz_dep_.CoverTab[187225]++
													return token, parts, &ValidationError{Inner: err, Errors: ValidationErrorMalformed}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:135
		// _ = "end of CoverTab[187225]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:136
		_go_fuzz_dep_.CoverTab[187226]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:136
		// _ = "end of CoverTab[187226]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:136
	// _ = "end of CoverTab[187207]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:136
	_go_fuzz_dep_.CoverTab[187208]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:139
	if method, ok := token.Header["alg"].(string); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:139
		_go_fuzz_dep_.CoverTab[187227]++
													if token.Method = GetSigningMethod(method); token.Method == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:140
			_go_fuzz_dep_.CoverTab[187228]++
														return token, parts, NewValidationError("signing method (alg) is unavailable.", ValidationErrorUnverifiable)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:141
			// _ = "end of CoverTab[187228]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:142
			_go_fuzz_dep_.CoverTab[187229]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:142
			// _ = "end of CoverTab[187229]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:142
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:142
		// _ = "end of CoverTab[187227]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:143
		_go_fuzz_dep_.CoverTab[187230]++
													return token, parts, NewValidationError("signing method (alg) is unspecified.", ValidationErrorUnverifiable)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:144
		// _ = "end of CoverTab[187230]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:145
	// _ = "end of CoverTab[187208]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:145
	_go_fuzz_dep_.CoverTab[187209]++

												return token, parts, nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:147
	// _ = "end of CoverTab[187209]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:148
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/parser.go:148
var _ = _go_fuzz_dep_.CoverTab
