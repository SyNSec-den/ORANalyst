// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:15
package auth

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:15
)

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	ecoidc "github.com/ericchiang/oidc"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"gopkg.in/square/go-jose.v2"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc/codes"

	"github.com/golang-jwt/jwt"
)

var log = logging.GetLogger("jwt")

const (
	// SharedSecretKey shared secret key for signing a token
	SharedSecretKey	= "SHARED_SECRET_KEY"
	// OIDCServerURL - will be accessed as Environment variable
	OIDCServerURL	= "OIDC_SERVER_URL"
	// OpenidConfiguration is the discovery point on the OIDC server
	OpenidConfiguration	= ".well-known/openid-configuration"
	// HS prefix for HS family algorithms
	HS	= "HS"
	// RS prefix for RS family algorithms
	RS	= "RS"
)

// JwtAuthenticator jwt authenticator
type JwtAuthenticator struct {
	publicKeys map[string][]byte
}

// ParseToken parse token and Ensure that the JWT conforms to the structure of a JWT.
func (j *JwtAuthenticator) parseToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:59
	_go_fuzz_dep_.CoverTab[190349]++
													claims := jwt.MapClaims{}
													token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:61
		_go_fuzz_dep_.CoverTab[190351]++

														if strings.HasPrefix(token.Method.Alg(), HS) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:63
			_go_fuzz_dep_.CoverTab[190353]++
															key := os.Getenv(SharedSecretKey)
															return []byte(key), nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:65
			// _ = "end of CoverTab[190353]"

		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:67
			_go_fuzz_dep_.CoverTab[190354]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:67
			if strings.HasPrefix(token.Method.Alg(), RS) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:67
				_go_fuzz_dep_.CoverTab[190355]++
																keyID, ok := token.Header["kid"]
																if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:69
					_go_fuzz_dep_.CoverTab[190359]++
																	return nil, status.Errorf(codes.Unauthenticated, "token header not found 'kid' (key ID)")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:70
					// _ = "end of CoverTab[190359]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:71
					_go_fuzz_dep_.CoverTab[190360]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:71
					// _ = "end of CoverTab[190360]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:71
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:71
				// _ = "end of CoverTab[190355]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:71
				_go_fuzz_dep_.CoverTab[190356]++
																keyIDStr := keyID.(string)
																publicKey, ok := j.publicKeys[keyIDStr]
																if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:74
					_go_fuzz_dep_.CoverTab[190361]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:77
					if err := j.refreshJwksKeys(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:77
						_go_fuzz_dep_.CoverTab[190363]++
																		return nil, status.Errorf(codes.Unauthenticated, "unable to refresh keys from ID provider %s", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:78
						// _ = "end of CoverTab[190363]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:79
						_go_fuzz_dep_.CoverTab[190364]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:79
						// _ = "end of CoverTab[190364]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:79
					}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:79
					// _ = "end of CoverTab[190361]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:79
					_go_fuzz_dep_.CoverTab[190362]++

																	if publicKey, ok = j.publicKeys[keyIDStr]; !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:81
						_go_fuzz_dep_.CoverTab[190365]++
																		return nil, status.Errorf(codes.Unauthenticated, "token has obsolete key ID %s", keyID)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:82
						// _ = "end of CoverTab[190365]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:83
						_go_fuzz_dep_.CoverTab[190366]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:83
						// _ = "end of CoverTab[190366]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:83
					}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:83
					// _ = "end of CoverTab[190362]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:84
					_go_fuzz_dep_.CoverTab[190367]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:84
					// _ = "end of CoverTab[190367]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:84
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:84
				// _ = "end of CoverTab[190356]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:84
				_go_fuzz_dep_.CoverTab[190357]++
																rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:86
					_go_fuzz_dep_.CoverTab[190368]++
																	return nil, status.Errorf(codes.Unauthenticated, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:87
					// _ = "end of CoverTab[190368]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:88
					_go_fuzz_dep_.CoverTab[190369]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:88
					// _ = "end of CoverTab[190369]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:88
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:88
				// _ = "end of CoverTab[190357]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:88
				_go_fuzz_dep_.CoverTab[190358]++
																return rsaPublicKey, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:89
				// _ = "end of CoverTab[190358]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:90
				_go_fuzz_dep_.CoverTab[190370]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:90
				// _ = "end of CoverTab[190370]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:90
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:90
			// _ = "end of CoverTab[190354]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:90
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:90
		// _ = "end of CoverTab[190351]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:90
		_go_fuzz_dep_.CoverTab[190352]++
														return nil, status.Errorf(codes.Unauthenticated, "unknown signing algorithm: %s", token.Method.Alg())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:91
		// _ = "end of CoverTab[190352]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:92
	// _ = "end of CoverTab[190349]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:92
	_go_fuzz_dep_.CoverTab[190350]++

													return token, claims, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:94
	// _ = "end of CoverTab[190350]"

}

// ParseAndValidate parse a jwt string token and validate it
func (j *JwtAuthenticator) ParseAndValidate(tokenString string) (jwt.MapClaims, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:99
	_go_fuzz_dep_.CoverTab[190371]++
													token, claims, err := j.parseToken(tokenString)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:101
		_go_fuzz_dep_.CoverTab[190374]++
														log.Warnf("Error parsing token: %s", tokenString)
														log.Warnf("Error %s", err.Error())
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:104
		// _ = "end of CoverTab[190374]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:105
		_go_fuzz_dep_.CoverTab[190375]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:105
		// _ = "end of CoverTab[190375]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:105
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:105
	// _ = "end of CoverTab[190371]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:105
	_go_fuzz_dep_.CoverTab[190372]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:108
	if !token.Valid {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:108
		_go_fuzz_dep_.CoverTab[190376]++
														return nil, status.Errorf(codes.Unauthenticated, "token is not valid %v", token)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:109
		// _ = "end of CoverTab[190376]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:110
		_go_fuzz_dep_.CoverTab[190377]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:110
		// _ = "end of CoverTab[190377]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:110
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:110
	// _ = "end of CoverTab[190372]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:110
	_go_fuzz_dep_.CoverTab[190373]++

													return claims, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:112
	// _ = "end of CoverTab[190373]"
}

// Connect back to the OpenIDConnect server to retrieve the keys
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:115
// They are rotated every 6 hours by default - we keep the keys in a cache
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:115
// It's a 2 step process
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:115
// 1) connect to $OIDCServerURL/.well-known/openid-configuration and retrieve the JSON payload
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:115
// 2) lookup the "keys" parameter and get keys from $OIDCServerURL/keys
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:115
// The keys are in a public key format and are converted to RSA Public Keys
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:121
func (j *JwtAuthenticator) refreshJwksKeys() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:121
	_go_fuzz_dep_.CoverTab[190378]++
													oidcURL := os.Getenv(OIDCServerURL)

													client := new(http.Client)
													resOpenIDConfig, err := client.Get(fmt.Sprintf("%s/%s", oidcURL, OpenidConfiguration))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:126
		_go_fuzz_dep_.CoverTab[190389]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:127
		// _ = "end of CoverTab[190389]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:128
		_go_fuzz_dep_.CoverTab[190390]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:128
		// _ = "end of CoverTab[190390]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:128
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:128
	// _ = "end of CoverTab[190378]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:128
	_go_fuzz_dep_.CoverTab[190379]++
													if resOpenIDConfig.Body != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:129
		_go_fuzz_dep_.CoverTab[190391]++
														defer resOpenIDConfig.Body.Close()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:130
		// _ = "end of CoverTab[190391]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:131
		_go_fuzz_dep_.CoverTab[190392]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:131
		// _ = "end of CoverTab[190392]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:131
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:131
	// _ = "end of CoverTab[190379]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:131
	_go_fuzz_dep_.CoverTab[190380]++
													openIDConfigBody, readErr := ioutil.ReadAll(resOpenIDConfig.Body)
													if readErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:133
		_go_fuzz_dep_.CoverTab[190393]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:134
		// _ = "end of CoverTab[190393]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:135
		_go_fuzz_dep_.CoverTab[190394]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:135
		// _ = "end of CoverTab[190394]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:135
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:135
	// _ = "end of CoverTab[190380]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:135
	_go_fuzz_dep_.CoverTab[190381]++
													var openIDprovider ecoidc.Provider
													jsonErr := json.Unmarshal(openIDConfigBody, &openIDprovider)
													if jsonErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:138
		_go_fuzz_dep_.CoverTab[190395]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:139
		// _ = "end of CoverTab[190395]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:140
		_go_fuzz_dep_.CoverTab[190396]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:140
		// _ = "end of CoverTab[190396]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:140
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:140
	// _ = "end of CoverTab[190381]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:140
	_go_fuzz_dep_.CoverTab[190382]++
													resOpenIDKeys, err := client.Get(openIDprovider.JWKSURL)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:142
		_go_fuzz_dep_.CoverTab[190397]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:143
		// _ = "end of CoverTab[190397]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:144
		_go_fuzz_dep_.CoverTab[190398]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:144
		// _ = "end of CoverTab[190398]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:144
	// _ = "end of CoverTab[190382]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:144
	_go_fuzz_dep_.CoverTab[190383]++
													if resOpenIDKeys.Body != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:145
		_go_fuzz_dep_.CoverTab[190399]++
														defer resOpenIDKeys.Body.Close()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:146
		// _ = "end of CoverTab[190399]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:147
		_go_fuzz_dep_.CoverTab[190400]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:147
		// _ = "end of CoverTab[190400]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:147
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:147
	// _ = "end of CoverTab[190383]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:147
	_go_fuzz_dep_.CoverTab[190384]++
													bodyOpenIDKeys, readErr := ioutil.ReadAll(resOpenIDKeys.Body)
													if readErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:149
		_go_fuzz_dep_.CoverTab[190401]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:150
		// _ = "end of CoverTab[190401]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:151
		_go_fuzz_dep_.CoverTab[190402]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:151
		// _ = "end of CoverTab[190402]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:151
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:151
	// _ = "end of CoverTab[190384]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:151
	_go_fuzz_dep_.CoverTab[190385]++
													var jsonWebKeySet jose.JSONWebKeySet
													if err := json.Unmarshal(bodyOpenIDKeys, &jsonWebKeySet); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:153
		_go_fuzz_dep_.CoverTab[190403]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:154
		// _ = "end of CoverTab[190403]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:155
		_go_fuzz_dep_.CoverTab[190404]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:155
		// _ = "end of CoverTab[190404]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:155
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:155
	// _ = "end of CoverTab[190385]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:155
	_go_fuzz_dep_.CoverTab[190386]++

													if j.publicKeys == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:157
		_go_fuzz_dep_.CoverTab[190405]++
														j.publicKeys = make(map[string][]byte)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:158
		// _ = "end of CoverTab[190405]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:159
		_go_fuzz_dep_.CoverTab[190406]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:159
		// _ = "end of CoverTab[190406]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:159
	// _ = "end of CoverTab[190386]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:159
	_go_fuzz_dep_.CoverTab[190387]++
													for _, key := range jsonWebKeySet.Keys {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:160
		_go_fuzz_dep_.CoverTab[190407]++
														data, err := x509.MarshalPKIXPublicKey(key.Key)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:162
			_go_fuzz_dep_.CoverTab[190409]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:163
			// _ = "end of CoverTab[190409]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:164
			_go_fuzz_dep_.CoverTab[190410]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:164
			// _ = "end of CoverTab[190410]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:164
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:164
		// _ = "end of CoverTab[190407]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:164
		_go_fuzz_dep_.CoverTab[190408]++
														block := pem.Block{
			Type:	"PUBLIC KEY",
			Bytes:	data,
		}
														pemBytes := pem.EncodeToMemory(&block)
														j.publicKeys[key.KeyID] = pemBytes
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:170
		// _ = "end of CoverTab[190408]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:171
	// _ = "end of CoverTab[190387]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:171
	_go_fuzz_dep_.CoverTab[190388]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:173
	// _ = "end of CoverTab[190388]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:174
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/auth/jwt.go:174
var _ = _go_fuzz_dep_.CoverTab
