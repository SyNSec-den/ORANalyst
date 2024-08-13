//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:17
)

import (
	"crypto/elliptic"
	"errors"
	"fmt"
)

// KeyAlgorithm represents a key management algorithm.
type KeyAlgorithm string

// SignatureAlgorithm represents a signature (or MAC) algorithm.
type SignatureAlgorithm string

// ContentEncryption represents a content encryption algorithm.
type ContentEncryption string

// CompressionAlgorithm represents an algorithm used for plaintext compression.
type CompressionAlgorithm string

var (
	// ErrCryptoFailure represents an error in cryptographic primitive. This
	// occurs when, for example, a message had an invalid authentication tag or
	// could not be decrypted.
	ErrCryptoFailure	= errors.New("square/go-jose: error in cryptographic primitive")

	// ErrUnsupportedAlgorithm indicates that a selected algorithm is not
	// supported. This occurs when trying to instantiate an encrypter for an
	// algorithm that is not yet implemented.
	ErrUnsupportedAlgorithm	= errors.New("square/go-jose: unknown/unsupported algorithm")

	// ErrUnsupportedKeyType indicates that the given key type/format is not
	// supported. This occurs when trying to instantiate an encrypter and passing
	// it a key of an unrecognized type or with unsupported parameters, such as
	// an RSA private key with more than two primes.
	ErrUnsupportedKeyType	= errors.New("square/go-jose: unsupported key type/format")

	// ErrNotSupported serialization of object is not supported. This occurs when
	// trying to compact-serialize an object which can't be represented in
	// compact form.
	ErrNotSupported	= errors.New("square/go-jose: compact serialization not supported for object")

	// ErrUnprotectedNonce indicates that while parsing a JWS or JWE object, a
	// nonce header parameter was included in an unprotected header object.
	ErrUnprotectedNonce	= errors.New("square/go-jose: Nonce parameter included in unprotected header")
)

// Key management algorithms
const (
	RSA1_5			= KeyAlgorithm("RSA1_5")		// RSA-PKCS1v1.5
	RSA_OAEP		= KeyAlgorithm("RSA-OAEP")		// RSA-OAEP-SHA1
	RSA_OAEP_256		= KeyAlgorithm("RSA-OAEP-256")		// RSA-OAEP-SHA256
	A128KW			= KeyAlgorithm("A128KW")		// AES key wrap (128)
	A192KW			= KeyAlgorithm("A192KW")		// AES key wrap (192)
	A256KW			= KeyAlgorithm("A256KW")		// AES key wrap (256)
	DIRECT			= KeyAlgorithm("dir")			// Direct encryption
	ECDH_ES			= KeyAlgorithm("ECDH-ES")		// ECDH-ES
	ECDH_ES_A128KW		= KeyAlgorithm("ECDH-ES+A128KW")	// ECDH-ES + AES key wrap (128)
	ECDH_ES_A192KW		= KeyAlgorithm("ECDH-ES+A192KW")	// ECDH-ES + AES key wrap (192)
	ECDH_ES_A256KW		= KeyAlgorithm("ECDH-ES+A256KW")	// ECDH-ES + AES key wrap (256)
	A128GCMKW		= KeyAlgorithm("A128GCMKW")		// AES-GCM key wrap (128)
	A192GCMKW		= KeyAlgorithm("A192GCMKW")		// AES-GCM key wrap (192)
	A256GCMKW		= KeyAlgorithm("A256GCMKW")		// AES-GCM key wrap (256)
	PBES2_HS256_A128KW	= KeyAlgorithm("PBES2-HS256+A128KW")	// PBES2 + HMAC-SHA256 + AES key wrap (128)
	PBES2_HS384_A192KW	= KeyAlgorithm("PBES2-HS384+A192KW")	// PBES2 + HMAC-SHA384 + AES key wrap (192)
	PBES2_HS512_A256KW	= KeyAlgorithm("PBES2-HS512+A256KW")	// PBES2 + HMAC-SHA512 + AES key wrap (256)
)

// Signature algorithms
const (
	HS256	= SignatureAlgorithm("HS256")	// HMAC using SHA-256
	HS384	= SignatureAlgorithm("HS384")	// HMAC using SHA-384
	HS512	= SignatureAlgorithm("HS512")	// HMAC using SHA-512
	RS256	= SignatureAlgorithm("RS256")	// RSASSA-PKCS-v1.5 using SHA-256
	RS384	= SignatureAlgorithm("RS384")	// RSASSA-PKCS-v1.5 using SHA-384
	RS512	= SignatureAlgorithm("RS512")	// RSASSA-PKCS-v1.5 using SHA-512
	ES256	= SignatureAlgorithm("ES256")	// ECDSA using P-256 and SHA-256
	ES384	= SignatureAlgorithm("ES384")	// ECDSA using P-384 and SHA-384
	ES512	= SignatureAlgorithm("ES512")	// ECDSA using P-521 and SHA-512
	PS256	= SignatureAlgorithm("PS256")	// RSASSA-PSS using SHA256 and MGF1-SHA256
	PS384	= SignatureAlgorithm("PS384")	// RSASSA-PSS using SHA384 and MGF1-SHA384
	PS512	= SignatureAlgorithm("PS512")	// RSASSA-PSS using SHA512 and MGF1-SHA512
)

// Content encryption algorithms
const (
	A128CBC_HS256	= ContentEncryption("A128CBC-HS256")	// AES-CBC + HMAC-SHA256 (128)
	A192CBC_HS384	= ContentEncryption("A192CBC-HS384")	// AES-CBC + HMAC-SHA384 (192)
	A256CBC_HS512	= ContentEncryption("A256CBC-HS512")	// AES-CBC + HMAC-SHA512 (256)
	A128GCM		= ContentEncryption("A128GCM")		// AES-GCM (128)
	A192GCM		= ContentEncryption("A192GCM")		// AES-GCM (192)
	A256GCM		= ContentEncryption("A256GCM")		// AES-GCM (256)
)

// Compression algorithms
const (
	NONE	= CompressionAlgorithm("")	// No compression
	DEFLATE	= CompressionAlgorithm("DEF")	// DEFLATE (RFC 1951)
)

// rawHeader represents the JOSE header for JWE/JWS objects (used for parsing).
type rawHeader struct {
	Alg	string			`json:"alg,omitempty"`
	Enc	ContentEncryption	`json:"enc,omitempty"`
	Zip	CompressionAlgorithm	`json:"zip,omitempty"`
	Crit	[]string		`json:"crit,omitempty"`
	Apu	*byteBuffer		`json:"apu,omitempty"`
	Apv	*byteBuffer		`json:"apv,omitempty"`
	Epk	*JsonWebKey		`json:"epk,omitempty"`
	Iv	*byteBuffer		`json:"iv,omitempty"`
	Tag	*byteBuffer		`json:"tag,omitempty"`
	Jwk	*JsonWebKey		`json:"jwk,omitempty"`
	Kid	string			`json:"kid,omitempty"`
	Nonce	string			`json:"nonce,omitempty"`
}

// JoseHeader represents the read-only JOSE header for JWE/JWS objects.
type JoseHeader struct {
	KeyID		string
	JsonWebKey	*JsonWebKey
	Algorithm	string
	Nonce		string
}

// sanitized produces a cleaned-up header object from the raw JSON.
func (parsed rawHeader) sanitized() JoseHeader {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:142
	_go_fuzz_dep_.CoverTab[186531]++
											return JoseHeader{
		KeyID:		parsed.Kid,
		JsonWebKey:	parsed.Jwk,
		Algorithm:	parsed.Alg,
		Nonce:		parsed.Nonce,
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:148
	// _ = "end of CoverTab[186531]"
}

// Merge headers from src into dst, giving precedence to headers from l.
func (dst *rawHeader) merge(src *rawHeader) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:152
	_go_fuzz_dep_.CoverTab[186532]++
											if src == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:153
		_go_fuzz_dep_.CoverTab[186546]++
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:154
		// _ = "end of CoverTab[186546]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:155
		_go_fuzz_dep_.CoverTab[186547]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:155
		// _ = "end of CoverTab[186547]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:155
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:155
	// _ = "end of CoverTab[186532]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:155
	_go_fuzz_dep_.CoverTab[186533]++

											if dst.Alg == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:157
		_go_fuzz_dep_.CoverTab[186548]++
												dst.Alg = src.Alg
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:158
		// _ = "end of CoverTab[186548]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:159
		_go_fuzz_dep_.CoverTab[186549]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:159
		// _ = "end of CoverTab[186549]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:159
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:159
	// _ = "end of CoverTab[186533]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:159
	_go_fuzz_dep_.CoverTab[186534]++
											if dst.Enc == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:160
		_go_fuzz_dep_.CoverTab[186550]++
												dst.Enc = src.Enc
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:161
		// _ = "end of CoverTab[186550]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:162
		_go_fuzz_dep_.CoverTab[186551]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:162
		// _ = "end of CoverTab[186551]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:162
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:162
	// _ = "end of CoverTab[186534]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:162
	_go_fuzz_dep_.CoverTab[186535]++
											if dst.Zip == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:163
		_go_fuzz_dep_.CoverTab[186552]++
												dst.Zip = src.Zip
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:164
		// _ = "end of CoverTab[186552]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:165
		_go_fuzz_dep_.CoverTab[186553]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:165
		// _ = "end of CoverTab[186553]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:165
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:165
	// _ = "end of CoverTab[186535]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:165
	_go_fuzz_dep_.CoverTab[186536]++
											if dst.Crit == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:166
		_go_fuzz_dep_.CoverTab[186554]++
												dst.Crit = src.Crit
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:167
		// _ = "end of CoverTab[186554]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:168
		_go_fuzz_dep_.CoverTab[186555]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:168
		// _ = "end of CoverTab[186555]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:168
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:168
	// _ = "end of CoverTab[186536]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:168
	_go_fuzz_dep_.CoverTab[186537]++
											if dst.Crit == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:169
		_go_fuzz_dep_.CoverTab[186556]++
												dst.Crit = src.Crit
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:170
		// _ = "end of CoverTab[186556]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:171
		_go_fuzz_dep_.CoverTab[186557]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:171
		// _ = "end of CoverTab[186557]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:171
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:171
	// _ = "end of CoverTab[186537]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:171
	_go_fuzz_dep_.CoverTab[186538]++
											if dst.Apu == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:172
		_go_fuzz_dep_.CoverTab[186558]++
												dst.Apu = src.Apu
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:173
		// _ = "end of CoverTab[186558]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:174
		_go_fuzz_dep_.CoverTab[186559]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:174
		// _ = "end of CoverTab[186559]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:174
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:174
	// _ = "end of CoverTab[186538]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:174
	_go_fuzz_dep_.CoverTab[186539]++
											if dst.Apv == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:175
		_go_fuzz_dep_.CoverTab[186560]++
												dst.Apv = src.Apv
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:176
		// _ = "end of CoverTab[186560]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:177
		_go_fuzz_dep_.CoverTab[186561]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:177
		// _ = "end of CoverTab[186561]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:177
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:177
	// _ = "end of CoverTab[186539]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:177
	_go_fuzz_dep_.CoverTab[186540]++
											if dst.Epk == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:178
		_go_fuzz_dep_.CoverTab[186562]++
												dst.Epk = src.Epk
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:179
		// _ = "end of CoverTab[186562]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:180
		_go_fuzz_dep_.CoverTab[186563]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:180
		// _ = "end of CoverTab[186563]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:180
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:180
	// _ = "end of CoverTab[186540]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:180
	_go_fuzz_dep_.CoverTab[186541]++
											if dst.Iv == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:181
		_go_fuzz_dep_.CoverTab[186564]++
												dst.Iv = src.Iv
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:182
		// _ = "end of CoverTab[186564]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:183
		_go_fuzz_dep_.CoverTab[186565]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:183
		// _ = "end of CoverTab[186565]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:183
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:183
	// _ = "end of CoverTab[186541]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:183
	_go_fuzz_dep_.CoverTab[186542]++
											if dst.Tag == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:184
		_go_fuzz_dep_.CoverTab[186566]++
												dst.Tag = src.Tag
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:185
		// _ = "end of CoverTab[186566]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:186
		_go_fuzz_dep_.CoverTab[186567]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:186
		// _ = "end of CoverTab[186567]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:186
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:186
	// _ = "end of CoverTab[186542]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:186
	_go_fuzz_dep_.CoverTab[186543]++
											if dst.Kid == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:187
		_go_fuzz_dep_.CoverTab[186568]++
												dst.Kid = src.Kid
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:188
		// _ = "end of CoverTab[186568]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:189
		_go_fuzz_dep_.CoverTab[186569]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:189
		// _ = "end of CoverTab[186569]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:189
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:189
	// _ = "end of CoverTab[186543]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:189
	_go_fuzz_dep_.CoverTab[186544]++
											if dst.Jwk == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:190
		_go_fuzz_dep_.CoverTab[186570]++
												dst.Jwk = src.Jwk
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:191
		// _ = "end of CoverTab[186570]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:192
		_go_fuzz_dep_.CoverTab[186571]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:192
		// _ = "end of CoverTab[186571]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:192
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:192
	// _ = "end of CoverTab[186544]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:192
	_go_fuzz_dep_.CoverTab[186545]++
											if dst.Nonce == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:193
		_go_fuzz_dep_.CoverTab[186572]++
												dst.Nonce = src.Nonce
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:194
		// _ = "end of CoverTab[186572]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:195
		_go_fuzz_dep_.CoverTab[186573]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:195
		// _ = "end of CoverTab[186573]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:195
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:195
	// _ = "end of CoverTab[186545]"
}

// Get JOSE name of curve
func curveName(crv elliptic.Curve) (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:199
	_go_fuzz_dep_.CoverTab[186574]++
											switch crv {
	case elliptic.P256():
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:201
		_go_fuzz_dep_.CoverTab[186575]++
												return "P-256", nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:202
		// _ = "end of CoverTab[186575]"
	case elliptic.P384():
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:203
		_go_fuzz_dep_.CoverTab[186576]++
												return "P-384", nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:204
		// _ = "end of CoverTab[186576]"
	case elliptic.P521():
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:205
		_go_fuzz_dep_.CoverTab[186577]++
												return "P-521", nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:206
		// _ = "end of CoverTab[186577]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:207
		_go_fuzz_dep_.CoverTab[186578]++
												return "", fmt.Errorf("square/go-jose: unsupported/unknown elliptic curve")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:208
		// _ = "end of CoverTab[186578]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:209
	// _ = "end of CoverTab[186574]"
}

// Get size of curve in bytes
func curveSize(crv elliptic.Curve) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:213
	_go_fuzz_dep_.CoverTab[186579]++
											bits := crv.Params().BitSize

											div := bits / 8
											mod := bits % 8

											if mod == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:219
		_go_fuzz_dep_.CoverTab[186581]++
												return div
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:220
		// _ = "end of CoverTab[186581]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:221
		_go_fuzz_dep_.CoverTab[186582]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:221
		// _ = "end of CoverTab[186582]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:221
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:221
	// _ = "end of CoverTab[186579]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:221
	_go_fuzz_dep_.CoverTab[186580]++

											return div + 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:223
	// _ = "end of CoverTab[186580]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:224
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/shared.go:224
var _ = _go_fuzz_dep_.CoverTab
