//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:17
)

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"reflect"
	"strings"

	"golang.org/x/crypto/ed25519"

	"gopkg.in/square/go-jose.v2/json"
)

// rawJSONWebKey represents a public or private key in JWK format, used for parsing/serializing.
type rawJSONWebKey struct {
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
	X5c		[]string	`json:"x5c,omitempty"`
	X5u		*url.URL	`json:"x5u,omitempty"`
	X5tSHA1		string		`json:"x5t,omitempty"`
	X5tSHA256	string		`json:"x5t#S256,omitempty"`
}

// JSONWebKey represents a public or private key in JWK format.
type JSONWebKey struct {
	// Cryptographic key, can be a symmetric or asymmetric key.
	Key	interface{}
	// Key identifier, parsed from `kid` header.
	KeyID	string
	// Key algorithm, parsed from `alg` header.
	Algorithm	string
	// Key use, parsed from `use` header.
	Use	string

	// X.509 certificate chain, parsed from `x5c` header.
	Certificates	[]*x509.Certificate
	// X.509 certificate URL, parsed from `x5u` header.
	CertificatesURL	*url.URL
	// X.509 certificate thumbprint (SHA-1), parsed from `x5t` header.
	CertificateThumbprintSHA1	[]byte
	// X.509 certificate thumbprint (SHA-256), parsed from `x5t#S256` header.
	CertificateThumbprintSHA256	[]byte
}

// MarshalJSON serializes the given key to its JSON representation.
func (k JSONWebKey) MarshalJSON() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:93
	_go_fuzz_dep_.CoverTab[189437]++
										var raw *rawJSONWebKey
										var err error

										switch key := k.Key.(type) {
	case ed25519.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:98
		_go_fuzz_dep_.CoverTab[189444]++
											raw = fromEdPublicKey(key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:99
		// _ = "end of CoverTab[189444]"
	case *ecdsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:100
		_go_fuzz_dep_.CoverTab[189445]++
											raw, err = fromEcPublicKey(key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:101
		// _ = "end of CoverTab[189445]"
	case *rsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:102
		_go_fuzz_dep_.CoverTab[189446]++
											raw = fromRsaPublicKey(key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:103
		// _ = "end of CoverTab[189446]"
	case ed25519.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:104
		_go_fuzz_dep_.CoverTab[189447]++
											raw, err = fromEdPrivateKey(key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:105
		// _ = "end of CoverTab[189447]"
	case *ecdsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:106
		_go_fuzz_dep_.CoverTab[189448]++
											raw, err = fromEcPrivateKey(key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:107
		// _ = "end of CoverTab[189448]"
	case *rsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:108
		_go_fuzz_dep_.CoverTab[189449]++
											raw, err = fromRsaPrivateKey(key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:109
		// _ = "end of CoverTab[189449]"
	case []byte:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:110
		_go_fuzz_dep_.CoverTab[189450]++
											raw, err = fromSymmetricKey(key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:111
		// _ = "end of CoverTab[189450]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:112
		_go_fuzz_dep_.CoverTab[189451]++
											return nil, fmt.Errorf("square/go-jose: unknown key type '%s'", reflect.TypeOf(key))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:113
		// _ = "end of CoverTab[189451]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:114
	// _ = "end of CoverTab[189437]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:114
	_go_fuzz_dep_.CoverTab[189438]++

										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:116
		_go_fuzz_dep_.CoverTab[189452]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:117
		// _ = "end of CoverTab[189452]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:118
		_go_fuzz_dep_.CoverTab[189453]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:118
		// _ = "end of CoverTab[189453]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:118
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:118
	// _ = "end of CoverTab[189438]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:118
	_go_fuzz_dep_.CoverTab[189439]++

										raw.Kid = k.KeyID
										raw.Alg = k.Algorithm
										raw.Use = k.Use

										for _, cert := range k.Certificates {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:124
		_go_fuzz_dep_.CoverTab[189454]++
											raw.X5c = append(raw.X5c, base64.StdEncoding.EncodeToString(cert.Raw))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:125
		// _ = "end of CoverTab[189454]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:126
	// _ = "end of CoverTab[189439]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:126
	_go_fuzz_dep_.CoverTab[189440]++

										x5tSHA1Len := len(k.CertificateThumbprintSHA1)
										x5tSHA256Len := len(k.CertificateThumbprintSHA256)
										if x5tSHA1Len > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:130
		_go_fuzz_dep_.CoverTab[189455]++
											if x5tSHA1Len != sha1.Size {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:131
			_go_fuzz_dep_.CoverTab[189457]++
												return nil, fmt.Errorf("square/go-jose: invalid SHA-1 thumbprint (must be %d bytes, not %d)", sha1.Size, x5tSHA1Len)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:132
			// _ = "end of CoverTab[189457]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:133
			_go_fuzz_dep_.CoverTab[189458]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:133
			// _ = "end of CoverTab[189458]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:133
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:133
		// _ = "end of CoverTab[189455]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:133
		_go_fuzz_dep_.CoverTab[189456]++
											raw.X5tSHA1 = base64.RawURLEncoding.EncodeToString(k.CertificateThumbprintSHA1)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:134
		// _ = "end of CoverTab[189456]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:135
		_go_fuzz_dep_.CoverTab[189459]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:135
		// _ = "end of CoverTab[189459]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:135
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:135
	// _ = "end of CoverTab[189440]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:135
	_go_fuzz_dep_.CoverTab[189441]++
										if x5tSHA256Len > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:136
		_go_fuzz_dep_.CoverTab[189460]++
											if x5tSHA256Len != sha256.Size {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:137
			_go_fuzz_dep_.CoverTab[189462]++
												return nil, fmt.Errorf("square/go-jose: invalid SHA-256 thumbprint (must be %d bytes, not %d)", sha256.Size, x5tSHA256Len)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:138
			// _ = "end of CoverTab[189462]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:139
			_go_fuzz_dep_.CoverTab[189463]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:139
			// _ = "end of CoverTab[189463]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:139
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:139
		// _ = "end of CoverTab[189460]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:139
		_go_fuzz_dep_.CoverTab[189461]++
											raw.X5tSHA256 = base64.RawURLEncoding.EncodeToString(k.CertificateThumbprintSHA256)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:140
		// _ = "end of CoverTab[189461]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:141
		_go_fuzz_dep_.CoverTab[189464]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:141
		// _ = "end of CoverTab[189464]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:141
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:141
	// _ = "end of CoverTab[189441]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:141
	_go_fuzz_dep_.CoverTab[189442]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:147
	if len(k.Certificates) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:147
		_go_fuzz_dep_.CoverTab[189465]++
											expectedSHA1 := sha1.Sum(k.Certificates[0].Raw)
											expectedSHA256 := sha256.Sum256(k.Certificates[0].Raw)

											if len(k.CertificateThumbprintSHA1) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:151
			_go_fuzz_dep_.CoverTab[189467]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:151
			return !bytes.Equal(k.CertificateThumbprintSHA1, expectedSHA1[:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:151
			// _ = "end of CoverTab[189467]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:151
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:151
			_go_fuzz_dep_.CoverTab[189468]++
												return nil, errors.New("square/go-jose: invalid SHA-1 thumbprint, does not match cert chain")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:152
			// _ = "end of CoverTab[189468]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:153
			_go_fuzz_dep_.CoverTab[189469]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:153
			// _ = "end of CoverTab[189469]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:153
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:153
		// _ = "end of CoverTab[189465]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:153
		_go_fuzz_dep_.CoverTab[189466]++
											if len(k.CertificateThumbprintSHA256) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:154
			_go_fuzz_dep_.CoverTab[189470]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:154
			return !bytes.Equal(k.CertificateThumbprintSHA256, expectedSHA256[:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:154
			// _ = "end of CoverTab[189470]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:154
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:154
			_go_fuzz_dep_.CoverTab[189471]++
												return nil, errors.New("square/go-jose: invalid or SHA-256 thumbprint, does not match cert chain")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:155
			// _ = "end of CoverTab[189471]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:156
			_go_fuzz_dep_.CoverTab[189472]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:156
			// _ = "end of CoverTab[189472]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:156
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:156
		// _ = "end of CoverTab[189466]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:157
		_go_fuzz_dep_.CoverTab[189473]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:157
		// _ = "end of CoverTab[189473]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:157
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:157
	// _ = "end of CoverTab[189442]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:157
	_go_fuzz_dep_.CoverTab[189443]++

										raw.X5u = k.CertificatesURL

										return json.Marshal(raw)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:161
	// _ = "end of CoverTab[189443]"
}

// UnmarshalJSON reads a key from its JSON representation.
func (k *JSONWebKey) UnmarshalJSON(data []byte) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:165
	_go_fuzz_dep_.CoverTab[189474]++
										var raw rawJSONWebKey
										err = json.Unmarshal(data, &raw)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:168
		_go_fuzz_dep_.CoverTab[189488]++
											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:169
		// _ = "end of CoverTab[189488]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:170
		_go_fuzz_dep_.CoverTab[189489]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:170
		// _ = "end of CoverTab[189489]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:170
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:170
	// _ = "end of CoverTab[189474]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:170
	_go_fuzz_dep_.CoverTab[189475]++

										certs, err := parseCertificateChain(raw.X5c)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:173
		_go_fuzz_dep_.CoverTab[189490]++
											return fmt.Errorf("square/go-jose: failed to unmarshal x5c field: %s", err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:174
		// _ = "end of CoverTab[189490]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:175
		_go_fuzz_dep_.CoverTab[189491]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:175
		// _ = "end of CoverTab[189491]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:175
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:175
	// _ = "end of CoverTab[189475]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:175
	_go_fuzz_dep_.CoverTab[189476]++

										var key interface{}
										var certPub interface{}
										var keyPub interface{}

										if len(certs) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:181
		_go_fuzz_dep_.CoverTab[189492]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:188
		certPub = certs[0].PublicKey
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:188
		// _ = "end of CoverTab[189492]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:189
		_go_fuzz_dep_.CoverTab[189493]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:189
		// _ = "end of CoverTab[189493]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:189
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:189
	// _ = "end of CoverTab[189476]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:189
	_go_fuzz_dep_.CoverTab[189477]++

										switch raw.Kty {
	case "EC":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:192
		_go_fuzz_dep_.CoverTab[189494]++
											if raw.D != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:193
			_go_fuzz_dep_.CoverTab[189500]++
												key, err = raw.ecPrivateKey()
												if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:195
				_go_fuzz_dep_.CoverTab[189501]++
													keyPub = key.(*ecdsa.PrivateKey).Public()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:196
				// _ = "end of CoverTab[189501]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:197
				_go_fuzz_dep_.CoverTab[189502]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:197
				// _ = "end of CoverTab[189502]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:197
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:197
			// _ = "end of CoverTab[189500]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:198
			_go_fuzz_dep_.CoverTab[189503]++
												key, err = raw.ecPublicKey()
												keyPub = key
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:200
			// _ = "end of CoverTab[189503]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:201
		// _ = "end of CoverTab[189494]"
	case "RSA":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:202
		_go_fuzz_dep_.CoverTab[189495]++
											if raw.D != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:203
			_go_fuzz_dep_.CoverTab[189504]++
												key, err = raw.rsaPrivateKey()
												if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:205
				_go_fuzz_dep_.CoverTab[189505]++
													keyPub = key.(*rsa.PrivateKey).Public()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:206
				// _ = "end of CoverTab[189505]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:207
				_go_fuzz_dep_.CoverTab[189506]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:207
				// _ = "end of CoverTab[189506]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:207
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:207
			// _ = "end of CoverTab[189504]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:208
			_go_fuzz_dep_.CoverTab[189507]++
												key, err = raw.rsaPublicKey()
												keyPub = key
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:210
			// _ = "end of CoverTab[189507]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:211
		// _ = "end of CoverTab[189495]"
	case "oct":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:212
		_go_fuzz_dep_.CoverTab[189496]++
											if certPub != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:213
			_go_fuzz_dep_.CoverTab[189508]++
												return errors.New("square/go-jose: invalid JWK, found 'oct' (symmetric) key with cert chain")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:214
			// _ = "end of CoverTab[189508]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:215
			_go_fuzz_dep_.CoverTab[189509]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:215
			// _ = "end of CoverTab[189509]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:215
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:215
		// _ = "end of CoverTab[189496]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:215
		_go_fuzz_dep_.CoverTab[189497]++
											key, err = raw.symmetricKey()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:216
		// _ = "end of CoverTab[189497]"
	case "OKP":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:217
		_go_fuzz_dep_.CoverTab[189498]++
											if raw.Crv == "Ed25519" && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:218
			_go_fuzz_dep_.CoverTab[189510]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:218
			return raw.X != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:218
			// _ = "end of CoverTab[189510]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:218
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:218
			_go_fuzz_dep_.CoverTab[189511]++
												if raw.D != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:219
				_go_fuzz_dep_.CoverTab[189512]++
													key, err = raw.edPrivateKey()
													if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:221
					_go_fuzz_dep_.CoverTab[189513]++
														keyPub = key.(ed25519.PrivateKey).Public()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:222
					// _ = "end of CoverTab[189513]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:223
					_go_fuzz_dep_.CoverTab[189514]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:223
					// _ = "end of CoverTab[189514]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:223
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:223
				// _ = "end of CoverTab[189512]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:224
				_go_fuzz_dep_.CoverTab[189515]++
													key, err = raw.edPublicKey()
													keyPub = key
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:226
				// _ = "end of CoverTab[189515]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:227
			// _ = "end of CoverTab[189511]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:228
			_go_fuzz_dep_.CoverTab[189516]++
												err = fmt.Errorf("square/go-jose: unknown curve %s'", raw.Crv)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:229
			// _ = "end of CoverTab[189516]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:230
		// _ = "end of CoverTab[189498]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:231
		_go_fuzz_dep_.CoverTab[189499]++
											err = fmt.Errorf("square/go-jose: unknown json web key type '%s'", raw.Kty)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:232
		// _ = "end of CoverTab[189499]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:233
	// _ = "end of CoverTab[189477]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:233
	_go_fuzz_dep_.CoverTab[189478]++

										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:235
		_go_fuzz_dep_.CoverTab[189517]++
											return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:236
		// _ = "end of CoverTab[189517]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:237
		_go_fuzz_dep_.CoverTab[189518]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:237
		// _ = "end of CoverTab[189518]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:237
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:237
	// _ = "end of CoverTab[189478]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:237
	_go_fuzz_dep_.CoverTab[189479]++

										if certPub != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:239
		_go_fuzz_dep_.CoverTab[189519]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:239
		return keyPub != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:239
		// _ = "end of CoverTab[189519]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:239
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:239
		_go_fuzz_dep_.CoverTab[189520]++
											if !reflect.DeepEqual(certPub, keyPub) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:240
			_go_fuzz_dep_.CoverTab[189521]++
												return errors.New("square/go-jose: invalid JWK, public keys in key and x5c fields to not match")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:241
			// _ = "end of CoverTab[189521]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:242
			_go_fuzz_dep_.CoverTab[189522]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:242
			// _ = "end of CoverTab[189522]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:242
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:242
		// _ = "end of CoverTab[189520]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:243
		_go_fuzz_dep_.CoverTab[189523]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:243
		// _ = "end of CoverTab[189523]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:243
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:243
	// _ = "end of CoverTab[189479]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:243
	_go_fuzz_dep_.CoverTab[189480]++

										*k = JSONWebKey{Key: key, KeyID: raw.Kid, Algorithm: raw.Alg, Use: raw.Use, Certificates: certs}

										k.CertificatesURL = raw.X5u

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:251
	x5tSHA1bytes, err := base64.RawURLEncoding.DecodeString(raw.X5tSHA1)
	if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:252
		_go_fuzz_dep_.CoverTab[189524]++
											return errors.New("square/go-jose: invalid JWK, x5t header has invalid encoding")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:253
		// _ = "end of CoverTab[189524]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:254
		_go_fuzz_dep_.CoverTab[189525]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:254
		// _ = "end of CoverTab[189525]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:254
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:254
	// _ = "end of CoverTab[189480]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:254
	_go_fuzz_dep_.CoverTab[189481]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:260
	if len(x5tSHA1bytes) == 2*sha1.Size {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:260
		_go_fuzz_dep_.CoverTab[189526]++
											hx, err := hex.DecodeString(string(x5tSHA1bytes))
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:262
			_go_fuzz_dep_.CoverTab[189528]++
												return fmt.Errorf("square/go-jose: invalid JWK, unable to hex decode x5t: %v", err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:263
			// _ = "end of CoverTab[189528]"

		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:265
			_go_fuzz_dep_.CoverTab[189529]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:265
			// _ = "end of CoverTab[189529]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:265
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:265
		// _ = "end of CoverTab[189526]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:265
		_go_fuzz_dep_.CoverTab[189527]++
											x5tSHA1bytes = hx
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:266
		// _ = "end of CoverTab[189527]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:267
		_go_fuzz_dep_.CoverTab[189530]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:267
		// _ = "end of CoverTab[189530]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:267
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:267
	// _ = "end of CoverTab[189481]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:267
	_go_fuzz_dep_.CoverTab[189482]++

										k.CertificateThumbprintSHA1 = x5tSHA1bytes

										x5tSHA256bytes, err := base64.RawURLEncoding.DecodeString(raw.X5tSHA256)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:272
		_go_fuzz_dep_.CoverTab[189531]++
											return errors.New("square/go-jose: invalid JWK, x5t#S256 header has invalid encoding")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:273
		// _ = "end of CoverTab[189531]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:274
		_go_fuzz_dep_.CoverTab[189532]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:274
		// _ = "end of CoverTab[189532]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:274
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:274
	// _ = "end of CoverTab[189482]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:274
	_go_fuzz_dep_.CoverTab[189483]++

										if len(x5tSHA256bytes) == 2*sha256.Size {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:276
		_go_fuzz_dep_.CoverTab[189533]++
											hx256, err := hex.DecodeString(string(x5tSHA256bytes))
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:278
			_go_fuzz_dep_.CoverTab[189535]++
												return fmt.Errorf("square/go-jose: invalid JWK, unable to hex decode x5t#S256: %v", err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:279
			// _ = "end of CoverTab[189535]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:280
			_go_fuzz_dep_.CoverTab[189536]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:280
			// _ = "end of CoverTab[189536]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:280
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:280
		// _ = "end of CoverTab[189533]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:280
		_go_fuzz_dep_.CoverTab[189534]++
											x5tSHA256bytes = hx256
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:281
		// _ = "end of CoverTab[189534]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:282
		_go_fuzz_dep_.CoverTab[189537]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:282
		// _ = "end of CoverTab[189537]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:282
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:282
	// _ = "end of CoverTab[189483]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:282
	_go_fuzz_dep_.CoverTab[189484]++

										k.CertificateThumbprintSHA256 = x5tSHA256bytes

										x5tSHA1Len := len(k.CertificateThumbprintSHA1)
										x5tSHA256Len := len(k.CertificateThumbprintSHA256)
										if x5tSHA1Len > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:288
		_go_fuzz_dep_.CoverTab[189538]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:288
		return x5tSHA1Len != sha1.Size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:288
		// _ = "end of CoverTab[189538]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:288
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:288
		_go_fuzz_dep_.CoverTab[189539]++
											return errors.New("square/go-jose: invalid JWK, x5t header is of incorrect size")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:289
		// _ = "end of CoverTab[189539]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:290
		_go_fuzz_dep_.CoverTab[189540]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:290
		// _ = "end of CoverTab[189540]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:290
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:290
	// _ = "end of CoverTab[189484]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:290
	_go_fuzz_dep_.CoverTab[189485]++
										if x5tSHA256Len > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:291
		_go_fuzz_dep_.CoverTab[189541]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:291
		return x5tSHA256Len != sha256.Size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:291
		// _ = "end of CoverTab[189541]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:291
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:291
		_go_fuzz_dep_.CoverTab[189542]++
											return errors.New("square/go-jose: invalid JWK, x5t#S256 header is of incorrect size")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:292
		// _ = "end of CoverTab[189542]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:293
		_go_fuzz_dep_.CoverTab[189543]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:293
		// _ = "end of CoverTab[189543]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:293
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:293
	// _ = "end of CoverTab[189485]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:293
	_go_fuzz_dep_.CoverTab[189486]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:296
	if len(k.Certificates) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:296
		_go_fuzz_dep_.CoverTab[189544]++
											leaf := k.Certificates[0]
											sha1sum := sha1.Sum(leaf.Raw)
											sha256sum := sha256.Sum256(leaf.Raw)

											if len(k.CertificateThumbprintSHA1) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:301
			_go_fuzz_dep_.CoverTab[189546]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:301
			return !bytes.Equal(sha1sum[:], k.CertificateThumbprintSHA1)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:301
			// _ = "end of CoverTab[189546]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:301
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:301
			_go_fuzz_dep_.CoverTab[189547]++
												return errors.New("square/go-jose: invalid JWK, x5c thumbprint does not match x5t value")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:302
			// _ = "end of CoverTab[189547]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:303
			_go_fuzz_dep_.CoverTab[189548]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:303
			// _ = "end of CoverTab[189548]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:303
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:303
		// _ = "end of CoverTab[189544]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:303
		_go_fuzz_dep_.CoverTab[189545]++

											if len(k.CertificateThumbprintSHA256) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:305
			_go_fuzz_dep_.CoverTab[189549]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:305
			return !bytes.Equal(sha256sum[:], k.CertificateThumbprintSHA256)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:305
			// _ = "end of CoverTab[189549]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:305
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:305
			_go_fuzz_dep_.CoverTab[189550]++
												return errors.New("square/go-jose: invalid JWK, x5c thumbprint does not match x5t#S256 value")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:306
			// _ = "end of CoverTab[189550]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:307
			_go_fuzz_dep_.CoverTab[189551]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:307
			// _ = "end of CoverTab[189551]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:307
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:307
		// _ = "end of CoverTab[189545]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:308
		_go_fuzz_dep_.CoverTab[189552]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:308
		// _ = "end of CoverTab[189552]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:308
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:308
	// _ = "end of CoverTab[189486]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:308
	_go_fuzz_dep_.CoverTab[189487]++

										return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:310
	// _ = "end of CoverTab[189487]"
}

// JSONWebKeySet represents a JWK Set object.
type JSONWebKeySet struct {
	Keys []JSONWebKey `json:"keys"`
}

// Key convenience method returns keys by key ID. Specification states
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:318
// that a JWK Set "SHOULD" use distinct key IDs, but allows for some
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:318
// cases where they are not distinct. Hence method returns a slice
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:318
// of JSONWebKeys.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:322
func (s *JSONWebKeySet) Key(kid string) []JSONWebKey {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:322
	_go_fuzz_dep_.CoverTab[189553]++
										var keys []JSONWebKey
										for _, key := range s.Keys {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:324
		_go_fuzz_dep_.CoverTab[189555]++
											if key.KeyID == kid {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:325
			_go_fuzz_dep_.CoverTab[189556]++
												keys = append(keys, key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:326
			// _ = "end of CoverTab[189556]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:327
			_go_fuzz_dep_.CoverTab[189557]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:327
			// _ = "end of CoverTab[189557]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:327
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:327
		// _ = "end of CoverTab[189555]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:328
	// _ = "end of CoverTab[189553]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:328
	_go_fuzz_dep_.CoverTab[189554]++

										return keys
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:330
	// _ = "end of CoverTab[189554]"
}

const rsaThumbprintTemplate = `{"e":"%s","kty":"RSA","n":"%s"}`
const ecThumbprintTemplate = `{"crv":"%s","kty":"EC","x":"%s","y":"%s"}`
const edThumbprintTemplate = `{"crv":"%s","kty":"OKP",x":"%s"}`

func ecThumbprintInput(curve elliptic.Curve, x, y *big.Int) (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:337
	_go_fuzz_dep_.CoverTab[189558]++
										coordLength := curveSize(curve)
										crv, err := curveName(curve)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:340
		_go_fuzz_dep_.CoverTab[189561]++
											return "", err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:341
		// _ = "end of CoverTab[189561]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:342
		_go_fuzz_dep_.CoverTab[189562]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:342
		// _ = "end of CoverTab[189562]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:342
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:342
	// _ = "end of CoverTab[189558]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:342
	_go_fuzz_dep_.CoverTab[189559]++

										if len(x.Bytes()) > coordLength || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:344
		_go_fuzz_dep_.CoverTab[189563]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:344
		return len(y.Bytes()) > coordLength
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:344
		// _ = "end of CoverTab[189563]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:344
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:344
		_go_fuzz_dep_.CoverTab[189564]++
											return "", errors.New("square/go-jose: invalid elliptic key (too large)")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:345
		// _ = "end of CoverTab[189564]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:346
		_go_fuzz_dep_.CoverTab[189565]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:346
		// _ = "end of CoverTab[189565]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:346
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:346
	// _ = "end of CoverTab[189559]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:346
	_go_fuzz_dep_.CoverTab[189560]++

										return fmt.Sprintf(ecThumbprintTemplate, crv,
		newFixedSizeBuffer(x.Bytes(), coordLength).base64(),
		newFixedSizeBuffer(y.Bytes(), coordLength).base64()), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:350
	// _ = "end of CoverTab[189560]"
}

func rsaThumbprintInput(n *big.Int, e int) (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:353
	_go_fuzz_dep_.CoverTab[189566]++
										return fmt.Sprintf(rsaThumbprintTemplate,
		newBufferFromInt(uint64(e)).base64(),
		newBuffer(n.Bytes()).base64()), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:356
	// _ = "end of CoverTab[189566]"
}

func edThumbprintInput(ed ed25519.PublicKey) (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:359
	_go_fuzz_dep_.CoverTab[189567]++
										crv := "Ed25519"
										if len(ed) > 32 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:361
		_go_fuzz_dep_.CoverTab[189569]++
											return "", errors.New("square/go-jose: invalid elliptic key (too large)")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:362
		// _ = "end of CoverTab[189569]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:363
		_go_fuzz_dep_.CoverTab[189570]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:363
		// _ = "end of CoverTab[189570]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:363
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:363
	// _ = "end of CoverTab[189567]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:363
	_go_fuzz_dep_.CoverTab[189568]++
										return fmt.Sprintf(edThumbprintTemplate, crv,
		newFixedSizeBuffer(ed, 32).base64()), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:365
	// _ = "end of CoverTab[189568]"
}

// Thumbprint computes the JWK Thumbprint of a key using the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:368
// indicated hash algorithm.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:370
func (k *JSONWebKey) Thumbprint(hash crypto.Hash) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:370
	_go_fuzz_dep_.CoverTab[189571]++
										var input string
										var err error
										switch key := k.Key.(type) {
	case ed25519.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:374
		_go_fuzz_dep_.CoverTab[189574]++
											input, err = edThumbprintInput(key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:375
		// _ = "end of CoverTab[189574]"
	case *ecdsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:376
		_go_fuzz_dep_.CoverTab[189575]++
											input, err = ecThumbprintInput(key.Curve, key.X, key.Y)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:377
		// _ = "end of CoverTab[189575]"
	case *ecdsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:378
		_go_fuzz_dep_.CoverTab[189576]++
											input, err = ecThumbprintInput(key.Curve, key.X, key.Y)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:379
		// _ = "end of CoverTab[189576]"
	case *rsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:380
		_go_fuzz_dep_.CoverTab[189577]++
											input, err = rsaThumbprintInput(key.N, key.E)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:381
		// _ = "end of CoverTab[189577]"
	case *rsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:382
		_go_fuzz_dep_.CoverTab[189578]++
											input, err = rsaThumbprintInput(key.N, key.E)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:383
		// _ = "end of CoverTab[189578]"
	case ed25519.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:384
		_go_fuzz_dep_.CoverTab[189579]++
											input, err = edThumbprintInput(ed25519.PublicKey(key[32:]))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:385
		// _ = "end of CoverTab[189579]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:386
		_go_fuzz_dep_.CoverTab[189580]++
											return nil, fmt.Errorf("square/go-jose: unknown key type '%s'", reflect.TypeOf(key))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:387
		// _ = "end of CoverTab[189580]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:388
	// _ = "end of CoverTab[189571]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:388
	_go_fuzz_dep_.CoverTab[189572]++

										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:390
		_go_fuzz_dep_.CoverTab[189581]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:391
		// _ = "end of CoverTab[189581]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:392
		_go_fuzz_dep_.CoverTab[189582]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:392
		// _ = "end of CoverTab[189582]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:392
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:392
	// _ = "end of CoverTab[189572]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:392
	_go_fuzz_dep_.CoverTab[189573]++

										h := hash.New()
										h.Write([]byte(input))
										return h.Sum(nil), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:396
	// _ = "end of CoverTab[189573]"
}

// IsPublic returns true if the JWK represents a public key (not symmetric, not private).
func (k *JSONWebKey) IsPublic() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:400
	_go_fuzz_dep_.CoverTab[189583]++
										switch k.Key.(type) {
	case *ecdsa.PublicKey, *rsa.PublicKey, ed25519.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:402
		_go_fuzz_dep_.CoverTab[189584]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:403
		// _ = "end of CoverTab[189584]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:404
		_go_fuzz_dep_.CoverTab[189585]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:405
		// _ = "end of CoverTab[189585]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:406
	// _ = "end of CoverTab[189583]"
}

// Public creates JSONWebKey with corresponding publik key if JWK represents asymmetric private key.
func (k *JSONWebKey) Public() JSONWebKey {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:410
	_go_fuzz_dep_.CoverTab[189586]++
										if k.IsPublic() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:411
		_go_fuzz_dep_.CoverTab[189589]++
											return *k
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:412
		// _ = "end of CoverTab[189589]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:413
		_go_fuzz_dep_.CoverTab[189590]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:413
		// _ = "end of CoverTab[189590]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:413
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:413
	// _ = "end of CoverTab[189586]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:413
	_go_fuzz_dep_.CoverTab[189587]++
										ret := *k
										switch key := k.Key.(type) {
	case *ecdsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:416
		_go_fuzz_dep_.CoverTab[189591]++
											ret.Key = key.Public()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:417
		// _ = "end of CoverTab[189591]"
	case *rsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:418
		_go_fuzz_dep_.CoverTab[189592]++
											ret.Key = key.Public()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:419
		// _ = "end of CoverTab[189592]"
	case ed25519.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:420
		_go_fuzz_dep_.CoverTab[189593]++
											ret.Key = key.Public()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:421
		// _ = "end of CoverTab[189593]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:422
		_go_fuzz_dep_.CoverTab[189594]++
											return JSONWebKey{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:423
		// _ = "end of CoverTab[189594]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:424
	// _ = "end of CoverTab[189587]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:424
	_go_fuzz_dep_.CoverTab[189588]++
										return ret
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:425
	// _ = "end of CoverTab[189588]"
}

// Valid checks that the key contains the expected parameters.
func (k *JSONWebKey) Valid() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:429
	_go_fuzz_dep_.CoverTab[189595]++
										if k.Key == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:430
		_go_fuzz_dep_.CoverTab[189598]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:431
		// _ = "end of CoverTab[189598]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:432
		_go_fuzz_dep_.CoverTab[189599]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:432
		// _ = "end of CoverTab[189599]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:432
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:432
	// _ = "end of CoverTab[189595]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:432
	_go_fuzz_dep_.CoverTab[189596]++
										switch key := k.Key.(type) {
	case *ecdsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:434
		_go_fuzz_dep_.CoverTab[189600]++
											if key.Curve == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:435
			_go_fuzz_dep_.CoverTab[189607]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:435
			return key.X == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:435
			// _ = "end of CoverTab[189607]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:435
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:435
			_go_fuzz_dep_.CoverTab[189608]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:435
			return key.Y == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:435
			// _ = "end of CoverTab[189608]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:435
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:435
			_go_fuzz_dep_.CoverTab[189609]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:436
			// _ = "end of CoverTab[189609]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:437
			_go_fuzz_dep_.CoverTab[189610]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:437
			// _ = "end of CoverTab[189610]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:437
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:437
		// _ = "end of CoverTab[189600]"
	case *ecdsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:438
		_go_fuzz_dep_.CoverTab[189601]++
											if key.Curve == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:439
			_go_fuzz_dep_.CoverTab[189611]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:439
			return key.X == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:439
			// _ = "end of CoverTab[189611]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:439
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:439
			_go_fuzz_dep_.CoverTab[189612]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:439
			return key.Y == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:439
			// _ = "end of CoverTab[189612]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:439
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:439
			_go_fuzz_dep_.CoverTab[189613]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:439
			return key.D == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:439
			// _ = "end of CoverTab[189613]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:439
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:439
			_go_fuzz_dep_.CoverTab[189614]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:440
			// _ = "end of CoverTab[189614]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:441
			_go_fuzz_dep_.CoverTab[189615]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:441
			// _ = "end of CoverTab[189615]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:441
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:441
		// _ = "end of CoverTab[189601]"
	case *rsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:442
		_go_fuzz_dep_.CoverTab[189602]++
											if key.N == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:443
			_go_fuzz_dep_.CoverTab[189616]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:443
			return key.E == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:443
			// _ = "end of CoverTab[189616]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:443
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:443
			_go_fuzz_dep_.CoverTab[189617]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:444
			// _ = "end of CoverTab[189617]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:445
			_go_fuzz_dep_.CoverTab[189618]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:445
			// _ = "end of CoverTab[189618]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:445
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:445
		// _ = "end of CoverTab[189602]"
	case *rsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:446
		_go_fuzz_dep_.CoverTab[189603]++
											if key.N == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:447
			_go_fuzz_dep_.CoverTab[189619]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:447
			return key.E == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:447
			// _ = "end of CoverTab[189619]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:447
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:447
			_go_fuzz_dep_.CoverTab[189620]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:447
			return key.D == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:447
			// _ = "end of CoverTab[189620]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:447
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:447
			_go_fuzz_dep_.CoverTab[189621]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:447
			return len(key.Primes) < 2
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:447
			// _ = "end of CoverTab[189621]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:447
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:447
			_go_fuzz_dep_.CoverTab[189622]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:448
			// _ = "end of CoverTab[189622]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:449
			_go_fuzz_dep_.CoverTab[189623]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:449
			// _ = "end of CoverTab[189623]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:449
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:449
		// _ = "end of CoverTab[189603]"
	case ed25519.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:450
		_go_fuzz_dep_.CoverTab[189604]++
											if len(key) != 32 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:451
			_go_fuzz_dep_.CoverTab[189624]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:452
			// _ = "end of CoverTab[189624]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:453
			_go_fuzz_dep_.CoverTab[189625]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:453
			// _ = "end of CoverTab[189625]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:453
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:453
		// _ = "end of CoverTab[189604]"
	case ed25519.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:454
		_go_fuzz_dep_.CoverTab[189605]++
											if len(key) != 64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:455
			_go_fuzz_dep_.CoverTab[189626]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:456
			// _ = "end of CoverTab[189626]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:457
			_go_fuzz_dep_.CoverTab[189627]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:457
			// _ = "end of CoverTab[189627]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:457
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:457
		// _ = "end of CoverTab[189605]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:458
		_go_fuzz_dep_.CoverTab[189606]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:459
		// _ = "end of CoverTab[189606]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:460
	// _ = "end of CoverTab[189596]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:460
	_go_fuzz_dep_.CoverTab[189597]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:461
	// _ = "end of CoverTab[189597]"
}

func (key rawJSONWebKey) rsaPublicKey() (*rsa.PublicKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:464
	_go_fuzz_dep_.CoverTab[189628]++
										if key.N == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:465
		_go_fuzz_dep_.CoverTab[189630]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:465
		return key.E == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:465
		// _ = "end of CoverTab[189630]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:465
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:465
		_go_fuzz_dep_.CoverTab[189631]++
											return nil, fmt.Errorf("square/go-jose: invalid RSA key, missing n/e values")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:466
		// _ = "end of CoverTab[189631]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:467
		_go_fuzz_dep_.CoverTab[189632]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:467
		// _ = "end of CoverTab[189632]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:467
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:467
	// _ = "end of CoverTab[189628]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:467
	_go_fuzz_dep_.CoverTab[189629]++

										return &rsa.PublicKey{
		N:	key.N.bigInt(),
		E:	key.E.toInt(),
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:472
	// _ = "end of CoverTab[189629]"
}

func fromEdPublicKey(pub ed25519.PublicKey) *rawJSONWebKey {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:475
	_go_fuzz_dep_.CoverTab[189633]++
										return &rawJSONWebKey{
		Kty:	"OKP",
		Crv:	"Ed25519",
		X:	newBuffer(pub),
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:480
	// _ = "end of CoverTab[189633]"
}

func fromRsaPublicKey(pub *rsa.PublicKey) *rawJSONWebKey {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:483
	_go_fuzz_dep_.CoverTab[189634]++
										return &rawJSONWebKey{
		Kty:	"RSA",
		N:	newBuffer(pub.N.Bytes()),
		E:	newBufferFromInt(uint64(pub.E)),
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:488
	// _ = "end of CoverTab[189634]"
}

func (key rawJSONWebKey) ecPublicKey() (*ecdsa.PublicKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:491
	_go_fuzz_dep_.CoverTab[189635]++
										var curve elliptic.Curve
										switch key.Crv {
	case "P-256":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:494
		_go_fuzz_dep_.CoverTab[189641]++
											curve = elliptic.P256()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:495
		// _ = "end of CoverTab[189641]"
	case "P-384":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:496
		_go_fuzz_dep_.CoverTab[189642]++
											curve = elliptic.P384()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:497
		// _ = "end of CoverTab[189642]"
	case "P-521":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:498
		_go_fuzz_dep_.CoverTab[189643]++
											curve = elliptic.P521()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:499
		// _ = "end of CoverTab[189643]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:500
		_go_fuzz_dep_.CoverTab[189644]++
											return nil, fmt.Errorf("square/go-jose: unsupported elliptic curve '%s'", key.Crv)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:501
		// _ = "end of CoverTab[189644]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:502
	// _ = "end of CoverTab[189635]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:502
	_go_fuzz_dep_.CoverTab[189636]++

										if key.X == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:504
		_go_fuzz_dep_.CoverTab[189645]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:504
		return key.Y == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:504
		// _ = "end of CoverTab[189645]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:504
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:504
		_go_fuzz_dep_.CoverTab[189646]++
											return nil, errors.New("square/go-jose: invalid EC key, missing x/y values")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:505
		// _ = "end of CoverTab[189646]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:506
		_go_fuzz_dep_.CoverTab[189647]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:506
		// _ = "end of CoverTab[189647]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:506
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:506
	// _ = "end of CoverTab[189636]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:506
	_go_fuzz_dep_.CoverTab[189637]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:511
	if curveSize(curve) != len(key.X.data) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:511
		_go_fuzz_dep_.CoverTab[189648]++
											return nil, fmt.Errorf("square/go-jose: invalid EC public key, wrong length for x")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:512
		// _ = "end of CoverTab[189648]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:513
		_go_fuzz_dep_.CoverTab[189649]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:513
		// _ = "end of CoverTab[189649]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:513
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:513
	// _ = "end of CoverTab[189637]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:513
	_go_fuzz_dep_.CoverTab[189638]++

										if curveSize(curve) != len(key.Y.data) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:515
		_go_fuzz_dep_.CoverTab[189650]++
											return nil, fmt.Errorf("square/go-jose: invalid EC public key, wrong length for y")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:516
		// _ = "end of CoverTab[189650]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:517
		_go_fuzz_dep_.CoverTab[189651]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:517
		// _ = "end of CoverTab[189651]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:517
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:517
	// _ = "end of CoverTab[189638]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:517
	_go_fuzz_dep_.CoverTab[189639]++

										x := key.X.bigInt()
										y := key.Y.bigInt()

										if !curve.IsOnCurve(x, y) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:522
		_go_fuzz_dep_.CoverTab[189652]++
											return nil, errors.New("square/go-jose: invalid EC key, X/Y are not on declared curve")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:523
		// _ = "end of CoverTab[189652]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:524
		_go_fuzz_dep_.CoverTab[189653]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:524
		// _ = "end of CoverTab[189653]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:524
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:524
	// _ = "end of CoverTab[189639]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:524
	_go_fuzz_dep_.CoverTab[189640]++

										return &ecdsa.PublicKey{
		Curve:	curve,
		X:	x,
		Y:	y,
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:530
	// _ = "end of CoverTab[189640]"
}

func fromEcPublicKey(pub *ecdsa.PublicKey) (*rawJSONWebKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:533
	_go_fuzz_dep_.CoverTab[189654]++
										if pub == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:534
		_go_fuzz_dep_.CoverTab[189658]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:534
		return pub.X == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:534
		// _ = "end of CoverTab[189658]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:534
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:534
		_go_fuzz_dep_.CoverTab[189659]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:534
		return pub.Y == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:534
		// _ = "end of CoverTab[189659]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:534
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:534
		_go_fuzz_dep_.CoverTab[189660]++
											return nil, fmt.Errorf("square/go-jose: invalid EC key (nil, or X/Y missing)")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:535
		// _ = "end of CoverTab[189660]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:536
		_go_fuzz_dep_.CoverTab[189661]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:536
		// _ = "end of CoverTab[189661]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:536
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:536
	// _ = "end of CoverTab[189654]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:536
	_go_fuzz_dep_.CoverTab[189655]++

										name, err := curveName(pub.Curve)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:539
		_go_fuzz_dep_.CoverTab[189662]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:540
		// _ = "end of CoverTab[189662]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:541
		_go_fuzz_dep_.CoverTab[189663]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:541
		// _ = "end of CoverTab[189663]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:541
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:541
	// _ = "end of CoverTab[189655]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:541
	_go_fuzz_dep_.CoverTab[189656]++

										size := curveSize(pub.Curve)

										xBytes := pub.X.Bytes()
										yBytes := pub.Y.Bytes()

										if len(xBytes) > size || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:548
		_go_fuzz_dep_.CoverTab[189664]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:548
		return len(yBytes) > size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:548
		// _ = "end of CoverTab[189664]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:548
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:548
		_go_fuzz_dep_.CoverTab[189665]++
											return nil, fmt.Errorf("square/go-jose: invalid EC key (X/Y too large)")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:549
		// _ = "end of CoverTab[189665]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:550
		_go_fuzz_dep_.CoverTab[189666]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:550
		// _ = "end of CoverTab[189666]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:550
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:550
	// _ = "end of CoverTab[189656]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:550
	_go_fuzz_dep_.CoverTab[189657]++

										key := &rawJSONWebKey{
		Kty:	"EC",
		Crv:	name,
		X:	newFixedSizeBuffer(xBytes, size),
		Y:	newFixedSizeBuffer(yBytes, size),
	}

										return key, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:559
	// _ = "end of CoverTab[189657]"
}

func (key rawJSONWebKey) edPrivateKey() (ed25519.PrivateKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:562
	_go_fuzz_dep_.CoverTab[189667]++
										var missing []string
										switch {
	case key.D == nil:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:565
		_go_fuzz_dep_.CoverTab[189670]++
											missing = append(missing, "D")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:566
		// _ = "end of CoverTab[189670]"
	case key.X == nil:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:567
		_go_fuzz_dep_.CoverTab[189671]++
											missing = append(missing, "X")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:568
		// _ = "end of CoverTab[189671]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:568
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:568
		_go_fuzz_dep_.CoverTab[189672]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:568
		// _ = "end of CoverTab[189672]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:569
	// _ = "end of CoverTab[189667]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:569
	_go_fuzz_dep_.CoverTab[189668]++

										if len(missing) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:571
		_go_fuzz_dep_.CoverTab[189673]++
											return nil, fmt.Errorf("square/go-jose: invalid Ed25519 private key, missing %s value(s)", strings.Join(missing, ", "))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:572
		// _ = "end of CoverTab[189673]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:573
		_go_fuzz_dep_.CoverTab[189674]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:573
		// _ = "end of CoverTab[189674]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:573
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:573
	// _ = "end of CoverTab[189668]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:573
	_go_fuzz_dep_.CoverTab[189669]++

										privateKey := make([]byte, ed25519.PrivateKeySize)
										copy(privateKey[0:32], key.D.bytes())
										copy(privateKey[32:], key.X.bytes())
										rv := ed25519.PrivateKey(privateKey)
										return rv, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:579
	// _ = "end of CoverTab[189669]"
}

func (key rawJSONWebKey) edPublicKey() (ed25519.PublicKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:582
	_go_fuzz_dep_.CoverTab[189675]++
										if key.X == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:583
		_go_fuzz_dep_.CoverTab[189677]++
											return nil, fmt.Errorf("square/go-jose: invalid Ed key, missing x value")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:584
		// _ = "end of CoverTab[189677]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:585
		_go_fuzz_dep_.CoverTab[189678]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:585
		// _ = "end of CoverTab[189678]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:585
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:585
	// _ = "end of CoverTab[189675]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:585
	_go_fuzz_dep_.CoverTab[189676]++
										publicKey := make([]byte, ed25519.PublicKeySize)
										copy(publicKey[0:32], key.X.bytes())
										rv := ed25519.PublicKey(publicKey)
										return rv, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:589
	// _ = "end of CoverTab[189676]"
}

func (key rawJSONWebKey) rsaPrivateKey() (*rsa.PrivateKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:592
	_go_fuzz_dep_.CoverTab[189679]++
										var missing []string
										switch {
	case key.N == nil:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:595
		_go_fuzz_dep_.CoverTab[189685]++
											missing = append(missing, "N")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:596
		// _ = "end of CoverTab[189685]"
	case key.E == nil:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:597
		_go_fuzz_dep_.CoverTab[189686]++
											missing = append(missing, "E")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:598
		// _ = "end of CoverTab[189686]"
	case key.D == nil:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:599
		_go_fuzz_dep_.CoverTab[189687]++
											missing = append(missing, "D")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:600
		// _ = "end of CoverTab[189687]"
	case key.P == nil:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:601
		_go_fuzz_dep_.CoverTab[189688]++
											missing = append(missing, "P")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:602
		// _ = "end of CoverTab[189688]"
	case key.Q == nil:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:603
		_go_fuzz_dep_.CoverTab[189689]++
											missing = append(missing, "Q")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:604
		// _ = "end of CoverTab[189689]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:604
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:604
		_go_fuzz_dep_.CoverTab[189690]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:604
		// _ = "end of CoverTab[189690]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:605
	// _ = "end of CoverTab[189679]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:605
	_go_fuzz_dep_.CoverTab[189680]++

										if len(missing) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:607
		_go_fuzz_dep_.CoverTab[189691]++
											return nil, fmt.Errorf("square/go-jose: invalid RSA private key, missing %s value(s)", strings.Join(missing, ", "))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:608
		// _ = "end of CoverTab[189691]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:609
		_go_fuzz_dep_.CoverTab[189692]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:609
		// _ = "end of CoverTab[189692]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:609
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:609
	// _ = "end of CoverTab[189680]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:609
	_go_fuzz_dep_.CoverTab[189681]++

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
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:623
		_go_fuzz_dep_.CoverTab[189693]++
											rv.Precomputed.Dp = key.Dp.bigInt()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:624
		// _ = "end of CoverTab[189693]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:625
		_go_fuzz_dep_.CoverTab[189694]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:625
		// _ = "end of CoverTab[189694]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:625
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:625
	// _ = "end of CoverTab[189681]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:625
	_go_fuzz_dep_.CoverTab[189682]++
										if key.Dq != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:626
		_go_fuzz_dep_.CoverTab[189695]++
											rv.Precomputed.Dq = key.Dq.bigInt()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:627
		// _ = "end of CoverTab[189695]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:628
		_go_fuzz_dep_.CoverTab[189696]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:628
		// _ = "end of CoverTab[189696]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:628
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:628
	// _ = "end of CoverTab[189682]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:628
	_go_fuzz_dep_.CoverTab[189683]++
										if key.Qi != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:629
		_go_fuzz_dep_.CoverTab[189697]++
											rv.Precomputed.Qinv = key.Qi.bigInt()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:630
		// _ = "end of CoverTab[189697]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:631
		_go_fuzz_dep_.CoverTab[189698]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:631
		// _ = "end of CoverTab[189698]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:631
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:631
	// _ = "end of CoverTab[189683]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:631
	_go_fuzz_dep_.CoverTab[189684]++

										err := rv.Validate()
										return rv, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:634
	// _ = "end of CoverTab[189684]"
}

func fromEdPrivateKey(ed ed25519.PrivateKey) (*rawJSONWebKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:637
	_go_fuzz_dep_.CoverTab[189699]++
										raw := fromEdPublicKey(ed25519.PublicKey(ed[32:]))

										raw.D = newBuffer(ed[0:32])
										return raw, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:641
	// _ = "end of CoverTab[189699]"
}

func fromRsaPrivateKey(rsa *rsa.PrivateKey) (*rawJSONWebKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:644
	_go_fuzz_dep_.CoverTab[189700]++
										if len(rsa.Primes) != 2 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:645
		_go_fuzz_dep_.CoverTab[189705]++
											return nil, ErrUnsupportedKeyType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:646
		// _ = "end of CoverTab[189705]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:647
		_go_fuzz_dep_.CoverTab[189706]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:647
		// _ = "end of CoverTab[189706]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:647
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:647
	// _ = "end of CoverTab[189700]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:647
	_go_fuzz_dep_.CoverTab[189701]++

										raw := fromRsaPublicKey(&rsa.PublicKey)

										raw.D = newBuffer(rsa.D.Bytes())
										raw.P = newBuffer(rsa.Primes[0].Bytes())
										raw.Q = newBuffer(rsa.Primes[1].Bytes())

										if rsa.Precomputed.Dp != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:655
		_go_fuzz_dep_.CoverTab[189707]++
											raw.Dp = newBuffer(rsa.Precomputed.Dp.Bytes())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:656
		// _ = "end of CoverTab[189707]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:657
		_go_fuzz_dep_.CoverTab[189708]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:657
		// _ = "end of CoverTab[189708]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:657
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:657
	// _ = "end of CoverTab[189701]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:657
	_go_fuzz_dep_.CoverTab[189702]++
										if rsa.Precomputed.Dq != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:658
		_go_fuzz_dep_.CoverTab[189709]++
											raw.Dq = newBuffer(rsa.Precomputed.Dq.Bytes())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:659
		// _ = "end of CoverTab[189709]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:660
		_go_fuzz_dep_.CoverTab[189710]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:660
		// _ = "end of CoverTab[189710]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:660
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:660
	// _ = "end of CoverTab[189702]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:660
	_go_fuzz_dep_.CoverTab[189703]++
										if rsa.Precomputed.Qinv != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:661
		_go_fuzz_dep_.CoverTab[189711]++
											raw.Qi = newBuffer(rsa.Precomputed.Qinv.Bytes())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:662
		// _ = "end of CoverTab[189711]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:663
		_go_fuzz_dep_.CoverTab[189712]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:663
		// _ = "end of CoverTab[189712]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:663
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:663
	// _ = "end of CoverTab[189703]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:663
	_go_fuzz_dep_.CoverTab[189704]++

										return raw, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:665
	// _ = "end of CoverTab[189704]"
}

func (key rawJSONWebKey) ecPrivateKey() (*ecdsa.PrivateKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:668
	_go_fuzz_dep_.CoverTab[189713]++
										var curve elliptic.Curve
										switch key.Crv {
	case "P-256":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:671
		_go_fuzz_dep_.CoverTab[189720]++
											curve = elliptic.P256()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:672
		// _ = "end of CoverTab[189720]"
	case "P-384":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:673
		_go_fuzz_dep_.CoverTab[189721]++
											curve = elliptic.P384()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:674
		// _ = "end of CoverTab[189721]"
	case "P-521":
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:675
		_go_fuzz_dep_.CoverTab[189722]++
											curve = elliptic.P521()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:676
		// _ = "end of CoverTab[189722]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:677
		_go_fuzz_dep_.CoverTab[189723]++
											return nil, fmt.Errorf("square/go-jose: unsupported elliptic curve '%s'", key.Crv)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:678
		// _ = "end of CoverTab[189723]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:679
	// _ = "end of CoverTab[189713]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:679
	_go_fuzz_dep_.CoverTab[189714]++

										if key.X == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:681
		_go_fuzz_dep_.CoverTab[189724]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:681
		return key.Y == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:681
		// _ = "end of CoverTab[189724]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:681
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:681
		_go_fuzz_dep_.CoverTab[189725]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:681
		return key.D == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:681
		// _ = "end of CoverTab[189725]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:681
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:681
		_go_fuzz_dep_.CoverTab[189726]++
											return nil, fmt.Errorf("square/go-jose: invalid EC private key, missing x/y/d values")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:682
		// _ = "end of CoverTab[189726]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:683
		_go_fuzz_dep_.CoverTab[189727]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:683
		// _ = "end of CoverTab[189727]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:683
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:683
	// _ = "end of CoverTab[189714]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:683
	_go_fuzz_dep_.CoverTab[189715]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:688
	if curveSize(curve) != len(key.X.data) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:688
		_go_fuzz_dep_.CoverTab[189728]++
											return nil, fmt.Errorf("square/go-jose: invalid EC private key, wrong length for x")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:689
		// _ = "end of CoverTab[189728]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:690
		_go_fuzz_dep_.CoverTab[189729]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:690
		// _ = "end of CoverTab[189729]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:690
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:690
	// _ = "end of CoverTab[189715]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:690
	_go_fuzz_dep_.CoverTab[189716]++

										if curveSize(curve) != len(key.Y.data) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:692
		_go_fuzz_dep_.CoverTab[189730]++
											return nil, fmt.Errorf("square/go-jose: invalid EC private key, wrong length for y")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:693
		// _ = "end of CoverTab[189730]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:694
		_go_fuzz_dep_.CoverTab[189731]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:694
		// _ = "end of CoverTab[189731]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:694
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:694
	// _ = "end of CoverTab[189716]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:694
	_go_fuzz_dep_.CoverTab[189717]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:697
	if dSize(curve) != len(key.D.data) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:697
		_go_fuzz_dep_.CoverTab[189732]++
											return nil, fmt.Errorf("square/go-jose: invalid EC private key, wrong length for d")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:698
		// _ = "end of CoverTab[189732]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:699
		_go_fuzz_dep_.CoverTab[189733]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:699
		// _ = "end of CoverTab[189733]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:699
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:699
	// _ = "end of CoverTab[189717]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:699
	_go_fuzz_dep_.CoverTab[189718]++

										x := key.X.bigInt()
										y := key.Y.bigInt()

										if !curve.IsOnCurve(x, y) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:704
		_go_fuzz_dep_.CoverTab[189734]++
											return nil, errors.New("square/go-jose: invalid EC key, X/Y are not on declared curve")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:705
		// _ = "end of CoverTab[189734]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:706
		_go_fuzz_dep_.CoverTab[189735]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:706
		// _ = "end of CoverTab[189735]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:706
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:706
	// _ = "end of CoverTab[189718]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:706
	_go_fuzz_dep_.CoverTab[189719]++

										return &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve:	curve,
			X:	x,
			Y:	y,
		},
		D:	key.D.bigInt(),
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:715
	// _ = "end of CoverTab[189719]"
}

func fromEcPrivateKey(ec *ecdsa.PrivateKey) (*rawJSONWebKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:718
	_go_fuzz_dep_.CoverTab[189736]++
										raw, err := fromEcPublicKey(&ec.PublicKey)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:720
		_go_fuzz_dep_.CoverTab[189739]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:721
		// _ = "end of CoverTab[189739]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:722
		_go_fuzz_dep_.CoverTab[189740]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:722
		// _ = "end of CoverTab[189740]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:722
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:722
	// _ = "end of CoverTab[189736]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:722
	_go_fuzz_dep_.CoverTab[189737]++

										if ec.D == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:724
		_go_fuzz_dep_.CoverTab[189741]++
											return nil, fmt.Errorf("square/go-jose: invalid EC private key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:725
		// _ = "end of CoverTab[189741]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:726
		_go_fuzz_dep_.CoverTab[189742]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:726
		// _ = "end of CoverTab[189742]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:726
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:726
	// _ = "end of CoverTab[189737]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:726
	_go_fuzz_dep_.CoverTab[189738]++

										raw.D = newFixedSizeBuffer(ec.D.Bytes(), dSize(ec.PublicKey.Curve))

										return raw, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:730
	// _ = "end of CoverTab[189738]"
}

// dSize returns the size in octets for the "d" member of an elliptic curve
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:733
// private key.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:733
// The length of this octet string MUST be ceiling(log-base-2(n)/8)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:733
// octets (where n is the order of the curve).
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:733
// https://tools.ietf.org/html/rfc7518#section-6.2.2.1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:738
func dSize(curve elliptic.Curve) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:738
	_go_fuzz_dep_.CoverTab[189743]++
										order := curve.Params().P
										bitLen := order.BitLen()
										size := bitLen / 8
										if bitLen%8 != 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:742
		_go_fuzz_dep_.CoverTab[189745]++
											size = size + 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:743
		// _ = "end of CoverTab[189745]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:744
		_go_fuzz_dep_.CoverTab[189746]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:744
		// _ = "end of CoverTab[189746]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:744
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:744
	// _ = "end of CoverTab[189743]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:744
	_go_fuzz_dep_.CoverTab[189744]++
										return size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:745
	// _ = "end of CoverTab[189744]"
}

func fromSymmetricKey(key []byte) (*rawJSONWebKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:748
	_go_fuzz_dep_.CoverTab[189747]++
										return &rawJSONWebKey{
		Kty:	"oct",
		K:	newBuffer(key),
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:752
	// _ = "end of CoverTab[189747]"
}

func (key rawJSONWebKey) symmetricKey() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:755
	_go_fuzz_dep_.CoverTab[189748]++
										if key.K == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:756
		_go_fuzz_dep_.CoverTab[189750]++
											return nil, fmt.Errorf("square/go-jose: invalid OCT (symmetric) key, missing k value")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:757
		// _ = "end of CoverTab[189750]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:758
		_go_fuzz_dep_.CoverTab[189751]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:758
		// _ = "end of CoverTab[189751]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:758
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:758
	// _ = "end of CoverTab[189748]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:758
	_go_fuzz_dep_.CoverTab[189749]++
										return key.K.bytes(), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:759
	// _ = "end of CoverTab[189749]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:760
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwk.go:760
var _ = _go_fuzz_dep_.CoverTab
