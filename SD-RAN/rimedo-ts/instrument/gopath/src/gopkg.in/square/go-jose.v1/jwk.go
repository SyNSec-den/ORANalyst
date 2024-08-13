//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:17
)

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"gopkg.in/square/go-jose.v1/json"
)

// rawJsonWebKey represents a public or private key in JWK format, used for parsing/serializing.
type rawJsonWebKey struct {
	Use	string		`json:"use,omitempty"`
	Kty	string		`json:"kty,omitempty"`
	Kid	string		`json:"kid,omitempty"`
	Crv	string		`json:"crv,omitempty"`
	Alg	string		`json:"alg,omitempty"`
	K	*byteBuffer	`json:"k,omitempty"`
	X	*byteBuffer	`json:"x,omitempty"`
	Y	*byteBuffer	`json:"y,omitempty"`
	N	*byteBuffer	`json:"n,omitempty"`
	E	*byteBuffer	`json:"e,omitempty"`
	// -- Following fields are only used for private keys --
	// RSA uses D, P and Q, while ECDSA uses only D. Fields Dp, Dq, and Qi are
	// completely optional. Therefore for RSA/ECDSA, D != nil is a contract that
	// we have a private key whereas D == nil means we have only a public key.
	D	*byteBuffer	`json:"d,omitempty"`
	P	*byteBuffer	`json:"p,omitempty"`
	Q	*byteBuffer	`json:"q,omitempty"`
	Dp	*byteBuffer	`json:"dp,omitempty"`
	Dq	*byteBuffer	`json:"dq,omitempty"`
	Qi	*byteBuffer	`json:"qi,omitempty"`
	// Certificates
	X5c	[]string	`json:"x5c,omitempty"`
}

// JsonWebKey represents a public or private key in JWK format.
type JsonWebKey struct {
	Key		interface{}
	Certificates	[]*x509.Certificate
	KeyID		string
	Algorithm	string
	Use		string
}

// MarshalJSON serializes the given key to its JSON representation.
func (k JsonWebKey) MarshalJSON() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:71
	_go_fuzz_dep_.CoverTab[186275]++
										var raw *rawJsonWebKey
										var err error

										switch key := k.Key.(type) {
	case *ecdsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:76
		_go_fuzz_dep_.CoverTab[186279]++
											raw, err = fromEcPublicKey(key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:77
		// _ = "end of CoverTab[186279]"
	case *rsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:78
		_go_fuzz_dep_.CoverTab[186280]++
											raw = fromRsaPublicKey(key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:79
		// _ = "end of CoverTab[186280]"
	case *ecdsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:80
		_go_fuzz_dep_.CoverTab[186281]++
											raw, err = fromEcPrivateKey(key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:81
		// _ = "end of CoverTab[186281]"
	case *rsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:82
		_go_fuzz_dep_.CoverTab[186282]++
											raw, err = fromRsaPrivateKey(key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:83
		// _ = "end of CoverTab[186282]"
	case []byte:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:84
		_go_fuzz_dep_.CoverTab[186283]++
											raw, err = fromSymmetricKey(key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:85
		// _ = "end of CoverTab[186283]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:86
		_go_fuzz_dep_.CoverTab[186284]++
											return nil, fmt.Errorf("square/go-jose: unknown key type '%s'", reflect.TypeOf(key))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:87
		// _ = "end of CoverTab[186284]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:88
	// _ = "end of CoverTab[186275]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:88
	_go_fuzz_dep_.CoverTab[186276]++

										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:90
		_go_fuzz_dep_.CoverTab[186285]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:91
		// _ = "end of CoverTab[186285]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:92
		_go_fuzz_dep_.CoverTab[186286]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:92
		// _ = "end of CoverTab[186286]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:92
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:92
	// _ = "end of CoverTab[186276]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:92
	_go_fuzz_dep_.CoverTab[186277]++

										raw.Kid = k.KeyID
										raw.Alg = k.Algorithm
										raw.Use = k.Use

										for _, cert := range k.Certificates {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:98
		_go_fuzz_dep_.CoverTab[186287]++
											raw.X5c = append(raw.X5c, base64.StdEncoding.EncodeToString(cert.Raw))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:99
		// _ = "end of CoverTab[186287]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:100
	// _ = "end of CoverTab[186277]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:100
	_go_fuzz_dep_.CoverTab[186278]++

										return json.Marshal(raw)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:102
	// _ = "end of CoverTab[186278]"
}

// UnmarshalJSON reads a key from its JSON representation.
func (k *JsonWebKey) UnmarshalJSON(data []byte) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:106
	_go_fuzz_dep_.CoverTab[186288]++
										var raw rawJsonWebKey
										err = json.Unmarshal(data, &raw)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:109
		_go_fuzz_dep_.CoverTab[186293]++
											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:110
		// _ = "end of CoverTab[186293]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:111
		_go_fuzz_dep_.CoverTab[186294]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:111
		// _ = "end of CoverTab[186294]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:111
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:111
	// _ = "end of CoverTab[186288]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:111
	_go_fuzz_dep_.CoverTab[186289]++

										var key interface{}
										switch raw.Kty {
	case "EC":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:115
		_go_fuzz_dep_.CoverTab[186295]++
											if raw.D != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:116
			_go_fuzz_dep_.CoverTab[186299]++
												key, err = raw.ecPrivateKey()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:117
			// _ = "end of CoverTab[186299]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:118
			_go_fuzz_dep_.CoverTab[186300]++
												key, err = raw.ecPublicKey()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:119
			// _ = "end of CoverTab[186300]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:120
		// _ = "end of CoverTab[186295]"
	case "RSA":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:121
		_go_fuzz_dep_.CoverTab[186296]++
											if raw.D != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:122
			_go_fuzz_dep_.CoverTab[186301]++
												key, err = raw.rsaPrivateKey()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:123
			// _ = "end of CoverTab[186301]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:124
			_go_fuzz_dep_.CoverTab[186302]++
												key, err = raw.rsaPublicKey()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:125
			// _ = "end of CoverTab[186302]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:126
		// _ = "end of CoverTab[186296]"
	case "oct":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:127
		_go_fuzz_dep_.CoverTab[186297]++
											key, err = raw.symmetricKey()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:128
		// _ = "end of CoverTab[186297]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:129
		_go_fuzz_dep_.CoverTab[186298]++
											err = fmt.Errorf("square/go-jose: unknown json web key type '%s'", raw.Kty)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:130
		// _ = "end of CoverTab[186298]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:131
	// _ = "end of CoverTab[186289]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:131
	_go_fuzz_dep_.CoverTab[186290]++

										if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:133
		_go_fuzz_dep_.CoverTab[186303]++
											*k = JsonWebKey{Key: key, KeyID: raw.Kid, Algorithm: raw.Alg, Use: raw.Use}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:134
		// _ = "end of CoverTab[186303]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:135
		_go_fuzz_dep_.CoverTab[186304]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:135
		// _ = "end of CoverTab[186304]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:135
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:135
	// _ = "end of CoverTab[186290]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:135
	_go_fuzz_dep_.CoverTab[186291]++

										k.Certificates = make([]*x509.Certificate, len(raw.X5c))
										for i, cert := range raw.X5c {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:138
		_go_fuzz_dep_.CoverTab[186305]++
											raw, err := base64.StdEncoding.DecodeString(cert)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:140
			_go_fuzz_dep_.CoverTab[186307]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:141
			// _ = "end of CoverTab[186307]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:142
			_go_fuzz_dep_.CoverTab[186308]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:142
			// _ = "end of CoverTab[186308]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:142
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:142
		// _ = "end of CoverTab[186305]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:142
		_go_fuzz_dep_.CoverTab[186306]++
											k.Certificates[i], err = x509.ParseCertificate(raw)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:144
			_go_fuzz_dep_.CoverTab[186309]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:145
			// _ = "end of CoverTab[186309]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:146
			_go_fuzz_dep_.CoverTab[186310]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:146
			// _ = "end of CoverTab[186310]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:146
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:146
		// _ = "end of CoverTab[186306]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:147
	// _ = "end of CoverTab[186291]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:147
	_go_fuzz_dep_.CoverTab[186292]++

										return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:149
	// _ = "end of CoverTab[186292]"
}

// JsonWebKeySet represents a JWK Set object.
type JsonWebKeySet struct {
	Keys []JsonWebKey `json:"keys"`
}

// Key convenience method returns keys by key ID. Specification states
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:157
// that a JWK Set "SHOULD" use distinct key IDs, but allows for some
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:157
// cases where they are not distinct. Hence method returns a slice
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:157
// of JsonWebKeys.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:161
func (s *JsonWebKeySet) Key(kid string) []JsonWebKey {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:161
	_go_fuzz_dep_.CoverTab[186311]++
										var keys []JsonWebKey
										for _, key := range s.Keys {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:163
		_go_fuzz_dep_.CoverTab[186313]++
											if key.KeyID == kid {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:164
			_go_fuzz_dep_.CoverTab[186314]++
												keys = append(keys, key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:165
			// _ = "end of CoverTab[186314]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:166
			_go_fuzz_dep_.CoverTab[186315]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:166
			// _ = "end of CoverTab[186315]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:166
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:166
		// _ = "end of CoverTab[186313]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:167
	// _ = "end of CoverTab[186311]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:167
	_go_fuzz_dep_.CoverTab[186312]++

										return keys
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:169
	// _ = "end of CoverTab[186312]"
}

const rsaThumbprintTemplate = `{"e":"%s","kty":"RSA","n":"%s"}`
const ecThumbprintTemplate = `{"crv":"%s","kty":"EC","x":"%s","y":"%s"}`

func ecThumbprintInput(curve elliptic.Curve, x, y *big.Int) (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:175
	_go_fuzz_dep_.CoverTab[186316]++
										coordLength := curveSize(curve)
										crv, err := curveName(curve)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:178
		_go_fuzz_dep_.CoverTab[186319]++
											return "", err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:179
		// _ = "end of CoverTab[186319]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:180
		_go_fuzz_dep_.CoverTab[186320]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:180
		// _ = "end of CoverTab[186320]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:180
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:180
	// _ = "end of CoverTab[186316]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:180
	_go_fuzz_dep_.CoverTab[186317]++

										if len(x.Bytes()) > coordLength || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:182
		_go_fuzz_dep_.CoverTab[186321]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:182
		return len(y.Bytes()) > coordLength
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:182
		// _ = "end of CoverTab[186321]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:182
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:182
		_go_fuzz_dep_.CoverTab[186322]++
											return "", errors.New("square/go-jose: invalid elliptic key (too large)")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:183
		// _ = "end of CoverTab[186322]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:184
		_go_fuzz_dep_.CoverTab[186323]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:184
		// _ = "end of CoverTab[186323]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:184
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:184
	// _ = "end of CoverTab[186317]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:184
	_go_fuzz_dep_.CoverTab[186318]++

										return fmt.Sprintf(ecThumbprintTemplate, crv,
		newFixedSizeBuffer(x.Bytes(), coordLength).base64(),
		newFixedSizeBuffer(y.Bytes(), coordLength).base64()), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:188
	// _ = "end of CoverTab[186318]"
}

func rsaThumbprintInput(n *big.Int, e int) (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:191
	_go_fuzz_dep_.CoverTab[186324]++
										return fmt.Sprintf(rsaThumbprintTemplate,
		newBufferFromInt(uint64(e)).base64(),
		newBuffer(n.Bytes()).base64()), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:194
	// _ = "end of CoverTab[186324]"
}

// Thumbprint computes the JWK Thumbprint of a key using the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:197
// indicated hash algorithm.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:199
func (k *JsonWebKey) Thumbprint(hash crypto.Hash) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:199
	_go_fuzz_dep_.CoverTab[186325]++
										var input string
										var err error
										switch key := k.Key.(type) {
	case *ecdsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:203
		_go_fuzz_dep_.CoverTab[186328]++
											input, err = ecThumbprintInput(key.Curve, key.X, key.Y)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:204
		// _ = "end of CoverTab[186328]"
	case *ecdsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:205
		_go_fuzz_dep_.CoverTab[186329]++
											input, err = ecThumbprintInput(key.Curve, key.X, key.Y)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:206
		// _ = "end of CoverTab[186329]"
	case *rsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:207
		_go_fuzz_dep_.CoverTab[186330]++
											input, err = rsaThumbprintInput(key.N, key.E)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:208
		// _ = "end of CoverTab[186330]"
	case *rsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:209
		_go_fuzz_dep_.CoverTab[186331]++
											input, err = rsaThumbprintInput(key.N, key.E)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:210
		// _ = "end of CoverTab[186331]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:211
		_go_fuzz_dep_.CoverTab[186332]++
											return nil, fmt.Errorf("square/go-jose: unknown key type '%s'", reflect.TypeOf(key))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:212
		// _ = "end of CoverTab[186332]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:213
	// _ = "end of CoverTab[186325]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:213
	_go_fuzz_dep_.CoverTab[186326]++

										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:215
		_go_fuzz_dep_.CoverTab[186333]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:216
		// _ = "end of CoverTab[186333]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:217
		_go_fuzz_dep_.CoverTab[186334]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:217
		// _ = "end of CoverTab[186334]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:217
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:217
	// _ = "end of CoverTab[186326]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:217
	_go_fuzz_dep_.CoverTab[186327]++

										h := hash.New()
										h.Write([]byte(input))
										return h.Sum(nil), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:221
	// _ = "end of CoverTab[186327]"
}

// IsPublic returns true if the JWK represents a public key (not symmetric, not private).
func (k *JsonWebKey) IsPublic() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:225
	_go_fuzz_dep_.CoverTab[186335]++
										switch k.Key.(type) {
	case *ecdsa.PublicKey, *rsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:227
		_go_fuzz_dep_.CoverTab[186336]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:228
		// _ = "end of CoverTab[186336]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:229
		_go_fuzz_dep_.CoverTab[186337]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:230
		// _ = "end of CoverTab[186337]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:231
	// _ = "end of CoverTab[186335]"
}

// Valid checks that the key contains the expected parameters.
func (k *JsonWebKey) Valid() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:235
	_go_fuzz_dep_.CoverTab[186338]++
										if k.Key == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:236
		_go_fuzz_dep_.CoverTab[186341]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:237
		// _ = "end of CoverTab[186341]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:238
		_go_fuzz_dep_.CoverTab[186342]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:238
		// _ = "end of CoverTab[186342]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:238
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:238
	// _ = "end of CoverTab[186338]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:238
	_go_fuzz_dep_.CoverTab[186339]++
										switch key := k.Key.(type) {
	case *ecdsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:240
		_go_fuzz_dep_.CoverTab[186343]++
											if key.Curve == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:241
			_go_fuzz_dep_.CoverTab[186348]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:241
			return key.X == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:241
			// _ = "end of CoverTab[186348]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:241
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:241
			_go_fuzz_dep_.CoverTab[186349]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:241
			return key.Y == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:241
			// _ = "end of CoverTab[186349]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:241
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:241
			_go_fuzz_dep_.CoverTab[186350]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:242
			// _ = "end of CoverTab[186350]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:243
			_go_fuzz_dep_.CoverTab[186351]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:243
			// _ = "end of CoverTab[186351]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:243
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:243
		// _ = "end of CoverTab[186343]"
	case *ecdsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:244
		_go_fuzz_dep_.CoverTab[186344]++
											if key.Curve == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:245
			_go_fuzz_dep_.CoverTab[186352]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:245
			return key.X == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:245
			// _ = "end of CoverTab[186352]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:245
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:245
			_go_fuzz_dep_.CoverTab[186353]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:245
			return key.Y == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:245
			// _ = "end of CoverTab[186353]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:245
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:245
			_go_fuzz_dep_.CoverTab[186354]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:245
			return key.D == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:245
			// _ = "end of CoverTab[186354]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:245
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:245
			_go_fuzz_dep_.CoverTab[186355]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:246
			// _ = "end of CoverTab[186355]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:247
			_go_fuzz_dep_.CoverTab[186356]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:247
			// _ = "end of CoverTab[186356]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:247
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:247
		// _ = "end of CoverTab[186344]"
	case *rsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:248
		_go_fuzz_dep_.CoverTab[186345]++
											if key.N == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:249
			_go_fuzz_dep_.CoverTab[186357]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:249
			return key.E == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:249
			// _ = "end of CoverTab[186357]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:249
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:249
			_go_fuzz_dep_.CoverTab[186358]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:250
			// _ = "end of CoverTab[186358]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:251
			_go_fuzz_dep_.CoverTab[186359]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:251
			// _ = "end of CoverTab[186359]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:251
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:251
		// _ = "end of CoverTab[186345]"
	case *rsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:252
		_go_fuzz_dep_.CoverTab[186346]++
											if key.N == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:253
			_go_fuzz_dep_.CoverTab[186360]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:253
			return key.E == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:253
			// _ = "end of CoverTab[186360]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:253
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:253
			_go_fuzz_dep_.CoverTab[186361]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:253
			return key.D == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:253
			// _ = "end of CoverTab[186361]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:253
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:253
			_go_fuzz_dep_.CoverTab[186362]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:253
			return len(key.Primes) < 2
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:253
			// _ = "end of CoverTab[186362]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:253
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:253
			_go_fuzz_dep_.CoverTab[186363]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:254
			// _ = "end of CoverTab[186363]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:255
			_go_fuzz_dep_.CoverTab[186364]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:255
			// _ = "end of CoverTab[186364]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:255
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:255
		// _ = "end of CoverTab[186346]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:256
		_go_fuzz_dep_.CoverTab[186347]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:257
		// _ = "end of CoverTab[186347]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:258
	// _ = "end of CoverTab[186339]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:258
	_go_fuzz_dep_.CoverTab[186340]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:259
	// _ = "end of CoverTab[186340]"
}

func (key rawJsonWebKey) rsaPublicKey() (*rsa.PublicKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:262
	_go_fuzz_dep_.CoverTab[186365]++
										if key.N == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:263
		_go_fuzz_dep_.CoverTab[186367]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:263
		return key.E == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:263
		// _ = "end of CoverTab[186367]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:263
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:263
		_go_fuzz_dep_.CoverTab[186368]++
											return nil, fmt.Errorf("square/go-jose: invalid RSA key, missing n/e values")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:264
		// _ = "end of CoverTab[186368]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:265
		_go_fuzz_dep_.CoverTab[186369]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:265
		// _ = "end of CoverTab[186369]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:265
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:265
	// _ = "end of CoverTab[186365]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:265
	_go_fuzz_dep_.CoverTab[186366]++

										return &rsa.PublicKey{
		N:	key.N.bigInt(),
		E:	key.E.toInt(),
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:270
	// _ = "end of CoverTab[186366]"
}

func fromRsaPublicKey(pub *rsa.PublicKey) *rawJsonWebKey {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:273
	_go_fuzz_dep_.CoverTab[186370]++
										return &rawJsonWebKey{
		Kty:	"RSA",
		N:	newBuffer(pub.N.Bytes()),
		E:	newBufferFromInt(uint64(pub.E)),
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:278
	// _ = "end of CoverTab[186370]"
}

func (key rawJsonWebKey) ecPublicKey() (*ecdsa.PublicKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:281
	_go_fuzz_dep_.CoverTab[186371]++
										var curve elliptic.Curve
										switch key.Crv {
	case "P-256":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:284
		_go_fuzz_dep_.CoverTab[186375]++
											curve = elliptic.P256()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:285
		// _ = "end of CoverTab[186375]"
	case "P-384":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:286
		_go_fuzz_dep_.CoverTab[186376]++
											curve = elliptic.P384()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:287
		// _ = "end of CoverTab[186376]"
	case "P-521":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:288
		_go_fuzz_dep_.CoverTab[186377]++
											curve = elliptic.P521()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:289
		// _ = "end of CoverTab[186377]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:290
		_go_fuzz_dep_.CoverTab[186378]++
											return nil, fmt.Errorf("square/go-jose: unsupported elliptic curve '%s'", key.Crv)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:291
		// _ = "end of CoverTab[186378]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:292
	// _ = "end of CoverTab[186371]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:292
	_go_fuzz_dep_.CoverTab[186372]++

										if key.X == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:294
		_go_fuzz_dep_.CoverTab[186379]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:294
		return key.Y == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:294
		// _ = "end of CoverTab[186379]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:294
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:294
		_go_fuzz_dep_.CoverTab[186380]++
											return nil, errors.New("square/go-jose: invalid EC key, missing x/y values")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:295
		// _ = "end of CoverTab[186380]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:296
		_go_fuzz_dep_.CoverTab[186381]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:296
		// _ = "end of CoverTab[186381]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:296
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:296
	// _ = "end of CoverTab[186372]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:296
	_go_fuzz_dep_.CoverTab[186373]++

										x := key.X.bigInt()
										y := key.Y.bigInt()

										if !curve.IsOnCurve(x, y) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:301
		_go_fuzz_dep_.CoverTab[186382]++
											return nil, errors.New("square/go-jose: invalid EC key, X/Y are not on declared curve")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:302
		// _ = "end of CoverTab[186382]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:303
		_go_fuzz_dep_.CoverTab[186383]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:303
		// _ = "end of CoverTab[186383]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:303
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:303
	// _ = "end of CoverTab[186373]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:303
	_go_fuzz_dep_.CoverTab[186374]++

										return &ecdsa.PublicKey{
		Curve:	curve,
		X:	x,
		Y:	y,
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:309
	// _ = "end of CoverTab[186374]"
}

func fromEcPublicKey(pub *ecdsa.PublicKey) (*rawJsonWebKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:312
	_go_fuzz_dep_.CoverTab[186384]++
										if pub == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:313
		_go_fuzz_dep_.CoverTab[186388]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:313
		return pub.X == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:313
		// _ = "end of CoverTab[186388]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:313
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:313
		_go_fuzz_dep_.CoverTab[186389]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:313
		return pub.Y == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:313
		// _ = "end of CoverTab[186389]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:313
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:313
		_go_fuzz_dep_.CoverTab[186390]++
											return nil, fmt.Errorf("square/go-jose: invalid EC key (nil, or X/Y missing)")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:314
		// _ = "end of CoverTab[186390]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:315
		_go_fuzz_dep_.CoverTab[186391]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:315
		// _ = "end of CoverTab[186391]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:315
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:315
	// _ = "end of CoverTab[186384]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:315
	_go_fuzz_dep_.CoverTab[186385]++

										name, err := curveName(pub.Curve)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:318
		_go_fuzz_dep_.CoverTab[186392]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:319
		// _ = "end of CoverTab[186392]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:320
		_go_fuzz_dep_.CoverTab[186393]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:320
		// _ = "end of CoverTab[186393]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:320
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:320
	// _ = "end of CoverTab[186385]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:320
	_go_fuzz_dep_.CoverTab[186386]++

										size := curveSize(pub.Curve)

										xBytes := pub.X.Bytes()
										yBytes := pub.Y.Bytes()

										if len(xBytes) > size || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:327
		_go_fuzz_dep_.CoverTab[186394]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:327
		return len(yBytes) > size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:327
		// _ = "end of CoverTab[186394]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:327
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:327
		_go_fuzz_dep_.CoverTab[186395]++
											return nil, fmt.Errorf("square/go-jose: invalid EC key (X/Y too large)")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:328
		// _ = "end of CoverTab[186395]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:329
		_go_fuzz_dep_.CoverTab[186396]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:329
		// _ = "end of CoverTab[186396]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:329
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:329
	// _ = "end of CoverTab[186386]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:329
	_go_fuzz_dep_.CoverTab[186387]++

										key := &rawJsonWebKey{
		Kty:	"EC",
		Crv:	name,
		X:	newFixedSizeBuffer(xBytes, size),
		Y:	newFixedSizeBuffer(yBytes, size),
	}

										return key, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:338
	// _ = "end of CoverTab[186387]"
}

func (key rawJsonWebKey) rsaPrivateKey() (*rsa.PrivateKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:341
	_go_fuzz_dep_.CoverTab[186397]++
										var missing []string
										switch {
	case key.N == nil:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:344
		_go_fuzz_dep_.CoverTab[186403]++
											missing = append(missing, "N")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:345
		// _ = "end of CoverTab[186403]"
	case key.E == nil:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:346
		_go_fuzz_dep_.CoverTab[186404]++
											missing = append(missing, "E")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:347
		// _ = "end of CoverTab[186404]"
	case key.D == nil:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:348
		_go_fuzz_dep_.CoverTab[186405]++
											missing = append(missing, "D")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:349
		// _ = "end of CoverTab[186405]"
	case key.P == nil:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:350
		_go_fuzz_dep_.CoverTab[186406]++
											missing = append(missing, "P")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:351
		// _ = "end of CoverTab[186406]"
	case key.Q == nil:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:352
		_go_fuzz_dep_.CoverTab[186407]++
											missing = append(missing, "Q")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:353
		// _ = "end of CoverTab[186407]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:353
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:353
		_go_fuzz_dep_.CoverTab[186408]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:353
		// _ = "end of CoverTab[186408]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:354
	// _ = "end of CoverTab[186397]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:354
	_go_fuzz_dep_.CoverTab[186398]++

										if len(missing) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:356
		_go_fuzz_dep_.CoverTab[186409]++
											return nil, fmt.Errorf("square/go-jose: invalid RSA private key, missing %s value(s)", strings.Join(missing, ", "))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:357
		// _ = "end of CoverTab[186409]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:358
		_go_fuzz_dep_.CoverTab[186410]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:358
		// _ = "end of CoverTab[186410]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:358
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:358
	// _ = "end of CoverTab[186398]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:358
	_go_fuzz_dep_.CoverTab[186399]++

										rv := &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N:	key.N.bigInt(),
			E:	key.E.toInt(),
		},
		D:	key.D.bigInt(),
		Primes: []*big.Int{
			key.P.bigInt(),
			key.Q.bigInt(),
		},
	}

	if key.Dp != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:372
		_go_fuzz_dep_.CoverTab[186411]++
											rv.Precomputed.Dp = key.Dp.bigInt()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:373
		// _ = "end of CoverTab[186411]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:374
		_go_fuzz_dep_.CoverTab[186412]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:374
		// _ = "end of CoverTab[186412]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:374
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:374
	// _ = "end of CoverTab[186399]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:374
	_go_fuzz_dep_.CoverTab[186400]++
										if key.Dq != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:375
		_go_fuzz_dep_.CoverTab[186413]++
											rv.Precomputed.Dq = key.Dq.bigInt()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:376
		// _ = "end of CoverTab[186413]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:377
		_go_fuzz_dep_.CoverTab[186414]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:377
		// _ = "end of CoverTab[186414]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:377
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:377
	// _ = "end of CoverTab[186400]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:377
	_go_fuzz_dep_.CoverTab[186401]++
										if key.Qi != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:378
		_go_fuzz_dep_.CoverTab[186415]++
											rv.Precomputed.Qinv = key.Qi.bigInt()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:379
		// _ = "end of CoverTab[186415]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:380
		_go_fuzz_dep_.CoverTab[186416]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:380
		// _ = "end of CoverTab[186416]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:380
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:380
	// _ = "end of CoverTab[186401]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:380
	_go_fuzz_dep_.CoverTab[186402]++

										err := rv.Validate()
										return rv, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:383
	// _ = "end of CoverTab[186402]"
}

func fromRsaPrivateKey(rsa *rsa.PrivateKey) (*rawJsonWebKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:386
	_go_fuzz_dep_.CoverTab[186417]++
										if len(rsa.Primes) != 2 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:387
		_go_fuzz_dep_.CoverTab[186419]++
											return nil, ErrUnsupportedKeyType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:388
		// _ = "end of CoverTab[186419]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:389
		_go_fuzz_dep_.CoverTab[186420]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:389
		// _ = "end of CoverTab[186420]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:389
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:389
	// _ = "end of CoverTab[186417]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:389
	_go_fuzz_dep_.CoverTab[186418]++

										raw := fromRsaPublicKey(&rsa.PublicKey)

										raw.D = newBuffer(rsa.D.Bytes())
										raw.P = newBuffer(rsa.Primes[0].Bytes())
										raw.Q = newBuffer(rsa.Primes[1].Bytes())

										return raw, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:397
	// _ = "end of CoverTab[186418]"
}

func (key rawJsonWebKey) ecPrivateKey() (*ecdsa.PrivateKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:400
	_go_fuzz_dep_.CoverTab[186421]++
										var curve elliptic.Curve
										switch key.Crv {
	case "P-256":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:403
		_go_fuzz_dep_.CoverTab[186425]++
											curve = elliptic.P256()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:404
		// _ = "end of CoverTab[186425]"
	case "P-384":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:405
		_go_fuzz_dep_.CoverTab[186426]++
											curve = elliptic.P384()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:406
		// _ = "end of CoverTab[186426]"
	case "P-521":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:407
		_go_fuzz_dep_.CoverTab[186427]++
											curve = elliptic.P521()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:408
		// _ = "end of CoverTab[186427]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:409
		_go_fuzz_dep_.CoverTab[186428]++
											return nil, fmt.Errorf("square/go-jose: unsupported elliptic curve '%s'", key.Crv)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:410
		// _ = "end of CoverTab[186428]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:411
	// _ = "end of CoverTab[186421]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:411
	_go_fuzz_dep_.CoverTab[186422]++

										if key.X == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:413
		_go_fuzz_dep_.CoverTab[186429]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:413
		return key.Y == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:413
		// _ = "end of CoverTab[186429]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:413
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:413
		_go_fuzz_dep_.CoverTab[186430]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:413
		return key.D == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:413
		// _ = "end of CoverTab[186430]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:413
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:413
		_go_fuzz_dep_.CoverTab[186431]++
											return nil, fmt.Errorf("square/go-jose: invalid EC private key, missing x/y/d values")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:414
		// _ = "end of CoverTab[186431]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:415
		_go_fuzz_dep_.CoverTab[186432]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:415
		// _ = "end of CoverTab[186432]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:415
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:415
	// _ = "end of CoverTab[186422]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:415
	_go_fuzz_dep_.CoverTab[186423]++

										x := key.X.bigInt()
										y := key.Y.bigInt()

										if !curve.IsOnCurve(x, y) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:420
		_go_fuzz_dep_.CoverTab[186433]++
											return nil, errors.New("square/go-jose: invalid EC key, X/Y are not on declared curve")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:421
		// _ = "end of CoverTab[186433]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:422
		_go_fuzz_dep_.CoverTab[186434]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:422
		// _ = "end of CoverTab[186434]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:422
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:422
	// _ = "end of CoverTab[186423]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:422
	_go_fuzz_dep_.CoverTab[186424]++

										return &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve:	curve,
			X:	x,
			Y:	y,
		},
		D:	key.D.bigInt(),
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:431
	// _ = "end of CoverTab[186424]"
}

func fromEcPrivateKey(ec *ecdsa.PrivateKey) (*rawJsonWebKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:434
	_go_fuzz_dep_.CoverTab[186435]++
										raw, err := fromEcPublicKey(&ec.PublicKey)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:436
		_go_fuzz_dep_.CoverTab[186438]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:437
		// _ = "end of CoverTab[186438]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:438
		_go_fuzz_dep_.CoverTab[186439]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:438
		// _ = "end of CoverTab[186439]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:438
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:438
	// _ = "end of CoverTab[186435]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:438
	_go_fuzz_dep_.CoverTab[186436]++

										if ec.D == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:440
		_go_fuzz_dep_.CoverTab[186440]++
											return nil, fmt.Errorf("square/go-jose: invalid EC private key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:441
		// _ = "end of CoverTab[186440]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:442
		_go_fuzz_dep_.CoverTab[186441]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:442
		// _ = "end of CoverTab[186441]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:442
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:442
	// _ = "end of CoverTab[186436]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:442
	_go_fuzz_dep_.CoverTab[186437]++

										raw.D = newBuffer(ec.D.Bytes())

										return raw, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:446
	// _ = "end of CoverTab[186437]"
}

func fromSymmetricKey(key []byte) (*rawJsonWebKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:449
	_go_fuzz_dep_.CoverTab[186442]++
										return &rawJsonWebKey{
		Kty:	"oct",
		K:	newBuffer(key),
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:453
	// _ = "end of CoverTab[186442]"
}

func (key rawJsonWebKey) symmetricKey() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:456
	_go_fuzz_dep_.CoverTab[186443]++
										if key.K == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:457
		_go_fuzz_dep_.CoverTab[186445]++
											return nil, fmt.Errorf("square/go-jose: invalid OCT (symmetric) key, missing k value")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:458
		// _ = "end of CoverTab[186445]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:459
		_go_fuzz_dep_.CoverTab[186446]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:459
		// _ = "end of CoverTab[186446]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:459
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:459
	// _ = "end of CoverTab[186443]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:459
	_go_fuzz_dep_.CoverTab[186444]++
										return key.K.bytes(), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:460
	// _ = "end of CoverTab[186444]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:461
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwk.go:461
var _ = _go_fuzz_dep_.CoverTab
