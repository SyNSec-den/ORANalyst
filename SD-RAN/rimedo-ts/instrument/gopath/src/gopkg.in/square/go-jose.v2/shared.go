//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:17
)

import (
	"crypto/elliptic"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"

	"gopkg.in/square/go-jose.v2/json"
)

// KeyAlgorithm represents a key management algorithm.
type KeyAlgorithm string

// SignatureAlgorithm represents a signature (or MAC) algorithm.
type SignatureAlgorithm string

// ContentEncryption represents a content encryption algorithm.
type ContentEncryption string

// CompressionAlgorithm represents an algorithm used for plaintext compression.
type CompressionAlgorithm string

// ContentType represents type of the contained data.
type ContentType string

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

	// ErrInvalidKeySize indicates that the given key is not the correct size
	// for the selected algorithm. This can occur, for example, when trying to
	// encrypt with AES-256 but passing only a 128-bit key as input.
	ErrInvalidKeySize	= errors.New("square/go-jose: invalid key size for algorithm")

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
	ED25519			= KeyAlgorithm("ED25519")
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
	EdDSA	= SignatureAlgorithm("EdDSA")
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

// A key in the protected header of a JWS object. Use of the Header...
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:131
// constants is preferred to enhance type safety.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:133
type HeaderKey string

const (
	HeaderType		HeaderKey	= "typ"	// string
	HeaderContentType			= "cty"	// string

	// These are set by go-jose and shouldn't need to be set by consumers of the
	// library.
	headerAlgorithm		= "alg"		// string
	headerEncryption	= "enc"		// ContentEncryption
	headerCompression	= "zip"		// CompressionAlgorithm
	headerCritical		= "crit"	// []string

	headerAPU	= "apu"	// *byteBuffer
	headerAPV	= "apv"	// *byteBuffer
	headerEPK	= "epk"	// *JSONWebKey
	headerIV	= "iv"	// *byteBuffer
	headerTag	= "tag"	// *byteBuffer
	headerX5c	= "x5c"	// []*x509.Certificate

	headerJWK	= "jwk"		// *JSONWebKey
	headerKeyID	= "kid"		// string
	headerNonce	= "nonce"	// string
	headerB64	= "b64"		// bool

	headerP2C	= "p2c"	// *byteBuffer (int)
	headerP2S	= "p2s"	// *byteBuffer ([]byte)

)

// supportedCritical is the set of supported extensions that are understood and processed.
var supportedCritical = map[string]bool{
	headerB64: true,
}

// rawHeader represents the JOSE header for JWE/JWS objects (used for parsing).
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:168
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:168
// The decoding of the constituent items is deferred because we want to marshal
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:168
// some members into particular structs rather than generic maps, but at the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:168
// same time we need to receive any extra fields unhandled by this library to
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:168
// pass through to consuming code in case it wants to examine them.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:174
type rawHeader map[HeaderKey]*json.RawMessage

// Header represents the read-only JOSE header for JWE/JWS objects.
type Header struct {
	KeyID		string
	JSONWebKey	*JSONWebKey
	Algorithm	string
	Nonce		string

	// Unverified certificate chain parsed from x5c header.
	certificates	[]*x509.Certificate

	// Any headers not recognised above get unmarshaled
	// from JSON in a generic manner and placed in this map.
	ExtraHeaders	map[HeaderKey]interface{}
}

// Certificates verifies & returns the certificate chain present
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:191
// in the x5c header field of a message, if one was present. Returns
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:191
// an error if there was no x5c header present or the chain could
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:191
// not be validated with the given verify options.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:195
func (h Header) Certificates(opts x509.VerifyOptions) ([][]*x509.Certificate, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:195
	_go_fuzz_dep_.CoverTab[189914]++
											if len(h.certificates) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:196
		_go_fuzz_dep_.CoverTab[189917]++
												return nil, errors.New("square/go-jose: no x5c header present in message")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:197
		// _ = "end of CoverTab[189917]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:198
		_go_fuzz_dep_.CoverTab[189918]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:198
		// _ = "end of CoverTab[189918]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:198
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:198
	// _ = "end of CoverTab[189914]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:198
	_go_fuzz_dep_.CoverTab[189915]++

											leaf := h.certificates[0]
											if opts.Intermediates == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:201
		_go_fuzz_dep_.CoverTab[189919]++
												opts.Intermediates = x509.NewCertPool()
												for _, intermediate := range h.certificates[1:] {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:203
			_go_fuzz_dep_.CoverTab[189920]++
													opts.Intermediates.AddCert(intermediate)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:204
			// _ = "end of CoverTab[189920]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:205
		// _ = "end of CoverTab[189919]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:206
		_go_fuzz_dep_.CoverTab[189921]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:206
		// _ = "end of CoverTab[189921]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:206
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:206
	// _ = "end of CoverTab[189915]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:206
	_go_fuzz_dep_.CoverTab[189916]++

											return leaf.Verify(opts)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:208
	// _ = "end of CoverTab[189916]"
}

func (parsed rawHeader) set(k HeaderKey, v interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:211
	_go_fuzz_dep_.CoverTab[189922]++
											b, err := json.Marshal(v)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:213
		_go_fuzz_dep_.CoverTab[189924]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:214
		// _ = "end of CoverTab[189924]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:215
		_go_fuzz_dep_.CoverTab[189925]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:215
		// _ = "end of CoverTab[189925]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:215
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:215
	// _ = "end of CoverTab[189922]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:215
	_go_fuzz_dep_.CoverTab[189923]++

											parsed[k] = makeRawMessage(b)
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:218
	// _ = "end of CoverTab[189923]"
}

// getString gets a string from the raw JSON, defaulting to "".
func (parsed rawHeader) getString(k HeaderKey) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:222
	_go_fuzz_dep_.CoverTab[189926]++
											v, ok := parsed[k]
											if !ok || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:224
		_go_fuzz_dep_.CoverTab[189929]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:224
		return v == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:224
		// _ = "end of CoverTab[189929]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:224
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:224
		_go_fuzz_dep_.CoverTab[189930]++
												return ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:225
		// _ = "end of CoverTab[189930]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:226
		_go_fuzz_dep_.CoverTab[189931]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:226
		// _ = "end of CoverTab[189931]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:226
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:226
	// _ = "end of CoverTab[189926]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:226
	_go_fuzz_dep_.CoverTab[189927]++
											var s string
											err := json.Unmarshal(*v, &s)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:229
		_go_fuzz_dep_.CoverTab[189932]++
												return ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:230
		// _ = "end of CoverTab[189932]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:231
		_go_fuzz_dep_.CoverTab[189933]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:231
		// _ = "end of CoverTab[189933]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:231
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:231
	// _ = "end of CoverTab[189927]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:231
	_go_fuzz_dep_.CoverTab[189928]++
											return s
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:232
	// _ = "end of CoverTab[189928]"
}

// getByteBuffer gets a byte buffer from the raw JSON. Returns (nil, nil) if
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:235
// not specified.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:237
func (parsed rawHeader) getByteBuffer(k HeaderKey) (*byteBuffer, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:237
	_go_fuzz_dep_.CoverTab[189934]++
											v := parsed[k]
											if v == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:239
		_go_fuzz_dep_.CoverTab[189937]++
												return nil, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:240
		// _ = "end of CoverTab[189937]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:241
		_go_fuzz_dep_.CoverTab[189938]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:241
		// _ = "end of CoverTab[189938]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:241
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:241
	// _ = "end of CoverTab[189934]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:241
	_go_fuzz_dep_.CoverTab[189935]++
											var bb *byteBuffer
											err := json.Unmarshal(*v, &bb)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:244
		_go_fuzz_dep_.CoverTab[189939]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:245
		// _ = "end of CoverTab[189939]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:246
		_go_fuzz_dep_.CoverTab[189940]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:246
		// _ = "end of CoverTab[189940]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:246
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:246
	// _ = "end of CoverTab[189935]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:246
	_go_fuzz_dep_.CoverTab[189936]++
											return bb, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:247
	// _ = "end of CoverTab[189936]"
}

// getAlgorithm extracts parsed "alg" from the raw JSON as a KeyAlgorithm.
func (parsed rawHeader) getAlgorithm() KeyAlgorithm {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:251
	_go_fuzz_dep_.CoverTab[189941]++
											return KeyAlgorithm(parsed.getString(headerAlgorithm))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:252
	// _ = "end of CoverTab[189941]"
}

// getSignatureAlgorithm extracts parsed "alg" from the raw JSON as a SignatureAlgorithm.
func (parsed rawHeader) getSignatureAlgorithm() SignatureAlgorithm {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:256
	_go_fuzz_dep_.CoverTab[189942]++
											return SignatureAlgorithm(parsed.getString(headerAlgorithm))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:257
	// _ = "end of CoverTab[189942]"
}

// getEncryption extracts parsed "enc" from the raw JSON.
func (parsed rawHeader) getEncryption() ContentEncryption {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:261
	_go_fuzz_dep_.CoverTab[189943]++
											return ContentEncryption(parsed.getString(headerEncryption))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:262
	// _ = "end of CoverTab[189943]"
}

// getCompression extracts parsed "zip" from the raw JSON.
func (parsed rawHeader) getCompression() CompressionAlgorithm {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:266
	_go_fuzz_dep_.CoverTab[189944]++
											return CompressionAlgorithm(parsed.getString(headerCompression))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:267
	// _ = "end of CoverTab[189944]"
}

func (parsed rawHeader) getNonce() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:270
	_go_fuzz_dep_.CoverTab[189945]++
											return parsed.getString(headerNonce)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:271
	// _ = "end of CoverTab[189945]"
}

// getEPK extracts parsed "epk" from the raw JSON.
func (parsed rawHeader) getEPK() (*JSONWebKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:275
	_go_fuzz_dep_.CoverTab[189946]++
											v := parsed[headerEPK]
											if v == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:277
		_go_fuzz_dep_.CoverTab[189949]++
												return nil, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:278
		// _ = "end of CoverTab[189949]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:279
		_go_fuzz_dep_.CoverTab[189950]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:279
		// _ = "end of CoverTab[189950]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:279
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:279
	// _ = "end of CoverTab[189946]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:279
	_go_fuzz_dep_.CoverTab[189947]++
											var epk *JSONWebKey
											err := json.Unmarshal(*v, &epk)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:282
		_go_fuzz_dep_.CoverTab[189951]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:283
		// _ = "end of CoverTab[189951]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:284
		_go_fuzz_dep_.CoverTab[189952]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:284
		// _ = "end of CoverTab[189952]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:284
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:284
	// _ = "end of CoverTab[189947]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:284
	_go_fuzz_dep_.CoverTab[189948]++
											return epk, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:285
	// _ = "end of CoverTab[189948]"
}

// getAPU extracts parsed "apu" from the raw JSON.
func (parsed rawHeader) getAPU() (*byteBuffer, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:289
	_go_fuzz_dep_.CoverTab[189953]++
											return parsed.getByteBuffer(headerAPU)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:290
	// _ = "end of CoverTab[189953]"
}

// getAPV extracts parsed "apv" from the raw JSON.
func (parsed rawHeader) getAPV() (*byteBuffer, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:294
	_go_fuzz_dep_.CoverTab[189954]++
											return parsed.getByteBuffer(headerAPV)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:295
	// _ = "end of CoverTab[189954]"
}

// getIV extracts parsed "iv" frpom the raw JSON.
func (parsed rawHeader) getIV() (*byteBuffer, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:299
	_go_fuzz_dep_.CoverTab[189955]++
											return parsed.getByteBuffer(headerIV)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:300
	// _ = "end of CoverTab[189955]"
}

// getTag extracts parsed "tag" frpom the raw JSON.
func (parsed rawHeader) getTag() (*byteBuffer, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:304
	_go_fuzz_dep_.CoverTab[189956]++
											return parsed.getByteBuffer(headerTag)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:305
	// _ = "end of CoverTab[189956]"
}

// getJWK extracts parsed "jwk" from the raw JSON.
func (parsed rawHeader) getJWK() (*JSONWebKey, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:309
	_go_fuzz_dep_.CoverTab[189957]++
											v := parsed[headerJWK]
											if v == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:311
		_go_fuzz_dep_.CoverTab[189960]++
												return nil, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:312
		// _ = "end of CoverTab[189960]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:313
		_go_fuzz_dep_.CoverTab[189961]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:313
		// _ = "end of CoverTab[189961]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:313
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:313
	// _ = "end of CoverTab[189957]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:313
	_go_fuzz_dep_.CoverTab[189958]++
											var jwk *JSONWebKey
											err := json.Unmarshal(*v, &jwk)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:316
		_go_fuzz_dep_.CoverTab[189962]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:317
		// _ = "end of CoverTab[189962]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:318
		_go_fuzz_dep_.CoverTab[189963]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:318
		// _ = "end of CoverTab[189963]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:318
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:318
	// _ = "end of CoverTab[189958]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:318
	_go_fuzz_dep_.CoverTab[189959]++
											return jwk, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:319
	// _ = "end of CoverTab[189959]"
}

// getCritical extracts parsed "crit" from the raw JSON. If omitted, it
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:322
// returns an empty slice.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:324
func (parsed rawHeader) getCritical() ([]string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:324
	_go_fuzz_dep_.CoverTab[189964]++
											v := parsed[headerCritical]
											if v == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:326
		_go_fuzz_dep_.CoverTab[189967]++
												return nil, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:327
		// _ = "end of CoverTab[189967]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:328
		_go_fuzz_dep_.CoverTab[189968]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:328
		// _ = "end of CoverTab[189968]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:328
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:328
	// _ = "end of CoverTab[189964]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:328
	_go_fuzz_dep_.CoverTab[189965]++

											var q []string
											err := json.Unmarshal(*v, &q)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:332
		_go_fuzz_dep_.CoverTab[189969]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:333
		// _ = "end of CoverTab[189969]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:334
		_go_fuzz_dep_.CoverTab[189970]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:334
		// _ = "end of CoverTab[189970]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:334
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:334
	// _ = "end of CoverTab[189965]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:334
	_go_fuzz_dep_.CoverTab[189966]++
											return q, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:335
	// _ = "end of CoverTab[189966]"
}

// getS2C extracts parsed "p2c" from the raw JSON.
func (parsed rawHeader) getP2C() (int, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:339
	_go_fuzz_dep_.CoverTab[189971]++
											v := parsed[headerP2C]
											if v == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:341
		_go_fuzz_dep_.CoverTab[189974]++
												return 0, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:342
		// _ = "end of CoverTab[189974]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:343
		_go_fuzz_dep_.CoverTab[189975]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:343
		// _ = "end of CoverTab[189975]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:343
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:343
	// _ = "end of CoverTab[189971]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:343
	_go_fuzz_dep_.CoverTab[189972]++

											var p2c int
											err := json.Unmarshal(*v, &p2c)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:347
		_go_fuzz_dep_.CoverTab[189976]++
												return 0, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:348
		// _ = "end of CoverTab[189976]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:349
		_go_fuzz_dep_.CoverTab[189977]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:349
		// _ = "end of CoverTab[189977]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:349
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:349
	// _ = "end of CoverTab[189972]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:349
	_go_fuzz_dep_.CoverTab[189973]++
											return p2c, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:350
	// _ = "end of CoverTab[189973]"
}

// getS2S extracts parsed "p2s" from the raw JSON.
func (parsed rawHeader) getP2S() (*byteBuffer, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:354
	_go_fuzz_dep_.CoverTab[189978]++
											return parsed.getByteBuffer(headerP2S)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:355
	// _ = "end of CoverTab[189978]"
}

// getB64 extracts parsed "b64" from the raw JSON, defaulting to true.
func (parsed rawHeader) getB64() (bool, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:359
	_go_fuzz_dep_.CoverTab[189979]++
											v := parsed[headerB64]
											if v == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:361
		_go_fuzz_dep_.CoverTab[189982]++
												return true, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:362
		// _ = "end of CoverTab[189982]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:363
		_go_fuzz_dep_.CoverTab[189983]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:363
		// _ = "end of CoverTab[189983]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:363
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:363
	// _ = "end of CoverTab[189979]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:363
	_go_fuzz_dep_.CoverTab[189980]++

											var b64 bool
											err := json.Unmarshal(*v, &b64)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:367
		_go_fuzz_dep_.CoverTab[189984]++
												return true, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:368
		// _ = "end of CoverTab[189984]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:369
		_go_fuzz_dep_.CoverTab[189985]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:369
		// _ = "end of CoverTab[189985]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:369
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:369
	// _ = "end of CoverTab[189980]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:369
	_go_fuzz_dep_.CoverTab[189981]++
											return b64, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:370
	// _ = "end of CoverTab[189981]"
}

// sanitized produces a cleaned-up header object from the raw JSON.
func (parsed rawHeader) sanitized() (h Header, err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:374
	_go_fuzz_dep_.CoverTab[189986]++
											for k, v := range parsed {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:375
		_go_fuzz_dep_.CoverTab[189988]++
												if v == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:376
			_go_fuzz_dep_.CoverTab[189990]++
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:377
			// _ = "end of CoverTab[189990]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:378
			_go_fuzz_dep_.CoverTab[189991]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:378
			// _ = "end of CoverTab[189991]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:378
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:378
		// _ = "end of CoverTab[189988]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:378
		_go_fuzz_dep_.CoverTab[189989]++
												switch k {
		case headerJWK:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:380
			_go_fuzz_dep_.CoverTab[189992]++
													var jwk *JSONWebKey
													err = json.Unmarshal(*v, &jwk)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:383
				_go_fuzz_dep_.CoverTab[190005]++
														err = fmt.Errorf("failed to unmarshal JWK: %v: %#v", err, string(*v))
														return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:385
				// _ = "end of CoverTab[190005]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:386
				_go_fuzz_dep_.CoverTab[190006]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:386
				// _ = "end of CoverTab[190006]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:386
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:386
			// _ = "end of CoverTab[189992]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:386
			_go_fuzz_dep_.CoverTab[189993]++
													h.JSONWebKey = jwk
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:387
			// _ = "end of CoverTab[189993]"
		case headerKeyID:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:388
			_go_fuzz_dep_.CoverTab[189994]++
													var s string
													err = json.Unmarshal(*v, &s)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:391
				_go_fuzz_dep_.CoverTab[190007]++
														err = fmt.Errorf("failed to unmarshal key ID: %v: %#v", err, string(*v))
														return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:393
				// _ = "end of CoverTab[190007]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:394
				_go_fuzz_dep_.CoverTab[190008]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:394
				// _ = "end of CoverTab[190008]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:394
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:394
			// _ = "end of CoverTab[189994]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:394
			_go_fuzz_dep_.CoverTab[189995]++
													h.KeyID = s
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:395
			// _ = "end of CoverTab[189995]"
		case headerAlgorithm:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:396
			_go_fuzz_dep_.CoverTab[189996]++
													var s string
													err = json.Unmarshal(*v, &s)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:399
				_go_fuzz_dep_.CoverTab[190009]++
														err = fmt.Errorf("failed to unmarshal algorithm: %v: %#v", err, string(*v))
														return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:401
				// _ = "end of CoverTab[190009]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:402
				_go_fuzz_dep_.CoverTab[190010]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:402
				// _ = "end of CoverTab[190010]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:402
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:402
			// _ = "end of CoverTab[189996]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:402
			_go_fuzz_dep_.CoverTab[189997]++
													h.Algorithm = s
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:403
			// _ = "end of CoverTab[189997]"
		case headerNonce:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:404
			_go_fuzz_dep_.CoverTab[189998]++
													var s string
													err = json.Unmarshal(*v, &s)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:407
				_go_fuzz_dep_.CoverTab[190011]++
														err = fmt.Errorf("failed to unmarshal nonce: %v: %#v", err, string(*v))
														return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:409
				// _ = "end of CoverTab[190011]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:410
				_go_fuzz_dep_.CoverTab[190012]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:410
				// _ = "end of CoverTab[190012]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:410
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:410
			// _ = "end of CoverTab[189998]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:410
			_go_fuzz_dep_.CoverTab[189999]++
													h.Nonce = s
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:411
			// _ = "end of CoverTab[189999]"
		case headerX5c:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:412
			_go_fuzz_dep_.CoverTab[190000]++
													c := []string{}
													err = json.Unmarshal(*v, &c)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:415
				_go_fuzz_dep_.CoverTab[190013]++
														err = fmt.Errorf("failed to unmarshal x5c header: %v: %#v", err, string(*v))
														return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:417
				// _ = "end of CoverTab[190013]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:418
				_go_fuzz_dep_.CoverTab[190014]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:418
				// _ = "end of CoverTab[190014]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:418
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:418
			// _ = "end of CoverTab[190000]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:418
			_go_fuzz_dep_.CoverTab[190001]++
													h.certificates, err = parseCertificateChain(c)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:420
				_go_fuzz_dep_.CoverTab[190015]++
														err = fmt.Errorf("failed to unmarshal x5c header: %v: %#v", err, string(*v))
														return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:422
				// _ = "end of CoverTab[190015]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:423
				_go_fuzz_dep_.CoverTab[190016]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:423
				// _ = "end of CoverTab[190016]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:423
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:423
			// _ = "end of CoverTab[190001]"
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:424
			_go_fuzz_dep_.CoverTab[190002]++
													if h.ExtraHeaders == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:425
				_go_fuzz_dep_.CoverTab[190017]++
														h.ExtraHeaders = map[HeaderKey]interface{}{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:426
				// _ = "end of CoverTab[190017]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:427
				_go_fuzz_dep_.CoverTab[190018]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:427
				// _ = "end of CoverTab[190018]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:427
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:427
			// _ = "end of CoverTab[190002]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:427
			_go_fuzz_dep_.CoverTab[190003]++
													var v2 interface{}
													err = json.Unmarshal(*v, &v2)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:430
				_go_fuzz_dep_.CoverTab[190019]++
														err = fmt.Errorf("failed to unmarshal value: %v: %#v", err, string(*v))
														return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:432
				// _ = "end of CoverTab[190019]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:433
				_go_fuzz_dep_.CoverTab[190020]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:433
				// _ = "end of CoverTab[190020]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:433
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:433
			// _ = "end of CoverTab[190003]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:433
			_go_fuzz_dep_.CoverTab[190004]++
													h.ExtraHeaders[k] = v2
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:434
			// _ = "end of CoverTab[190004]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:435
		// _ = "end of CoverTab[189989]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:436
	// _ = "end of CoverTab[189986]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:436
	_go_fuzz_dep_.CoverTab[189987]++
											return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:437
	// _ = "end of CoverTab[189987]"
}

func parseCertificateChain(chain []string) ([]*x509.Certificate, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:440
	_go_fuzz_dep_.CoverTab[190021]++
											out := make([]*x509.Certificate, len(chain))
											for i, cert := range chain {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:442
		_go_fuzz_dep_.CoverTab[190023]++
												raw, err := base64.StdEncoding.DecodeString(cert)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:444
			_go_fuzz_dep_.CoverTab[190025]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:445
			// _ = "end of CoverTab[190025]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:446
			_go_fuzz_dep_.CoverTab[190026]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:446
			// _ = "end of CoverTab[190026]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:446
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:446
		// _ = "end of CoverTab[190023]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:446
		_go_fuzz_dep_.CoverTab[190024]++
												out[i], err = x509.ParseCertificate(raw)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:448
			_go_fuzz_dep_.CoverTab[190027]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:449
			// _ = "end of CoverTab[190027]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:450
			_go_fuzz_dep_.CoverTab[190028]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:450
			// _ = "end of CoverTab[190028]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:450
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:450
		// _ = "end of CoverTab[190024]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:451
	// _ = "end of CoverTab[190021]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:451
	_go_fuzz_dep_.CoverTab[190022]++
											return out, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:452
	// _ = "end of CoverTab[190022]"
}

func (dst rawHeader) isSet(k HeaderKey) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:455
	_go_fuzz_dep_.CoverTab[190029]++
											dvr := dst[k]
											if dvr == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:457
		_go_fuzz_dep_.CoverTab[190033]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:458
		// _ = "end of CoverTab[190033]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:459
		_go_fuzz_dep_.CoverTab[190034]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:459
		// _ = "end of CoverTab[190034]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:459
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:459
	// _ = "end of CoverTab[190029]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:459
	_go_fuzz_dep_.CoverTab[190030]++

											var dv interface{}
											err := json.Unmarshal(*dvr, &dv)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:463
		_go_fuzz_dep_.CoverTab[190035]++
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:464
		// _ = "end of CoverTab[190035]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:465
		_go_fuzz_dep_.CoverTab[190036]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:465
		// _ = "end of CoverTab[190036]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:465
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:465
	// _ = "end of CoverTab[190030]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:465
	_go_fuzz_dep_.CoverTab[190031]++

											if dvStr, ok := dv.(string); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:467
		_go_fuzz_dep_.CoverTab[190037]++
												return dvStr != ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:468
		// _ = "end of CoverTab[190037]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:469
		_go_fuzz_dep_.CoverTab[190038]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:469
		// _ = "end of CoverTab[190038]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:469
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:469
	// _ = "end of CoverTab[190031]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:469
	_go_fuzz_dep_.CoverTab[190032]++

											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:471
	// _ = "end of CoverTab[190032]"
}

// Merge headers from src into dst, giving precedence to headers from l.
func (dst rawHeader) merge(src *rawHeader) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:475
	_go_fuzz_dep_.CoverTab[190039]++
											if src == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:476
		_go_fuzz_dep_.CoverTab[190041]++
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:477
		// _ = "end of CoverTab[190041]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:478
		_go_fuzz_dep_.CoverTab[190042]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:478
		// _ = "end of CoverTab[190042]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:478
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:478
	// _ = "end of CoverTab[190039]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:478
	_go_fuzz_dep_.CoverTab[190040]++

											for k, v := range *src {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:480
		_go_fuzz_dep_.CoverTab[190043]++
												if dst.isSet(k) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:481
			_go_fuzz_dep_.CoverTab[190045]++
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:482
			// _ = "end of CoverTab[190045]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:483
			_go_fuzz_dep_.CoverTab[190046]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:483
			// _ = "end of CoverTab[190046]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:483
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:483
		// _ = "end of CoverTab[190043]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:483
		_go_fuzz_dep_.CoverTab[190044]++

												dst[k] = v
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:485
		// _ = "end of CoverTab[190044]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:486
	// _ = "end of CoverTab[190040]"
}

// Get JOSE name of curve
func curveName(crv elliptic.Curve) (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:490
	_go_fuzz_dep_.CoverTab[190047]++
											switch crv {
	case elliptic.P256():
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:492
		_go_fuzz_dep_.CoverTab[190048]++
												return "P-256", nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:493
		// _ = "end of CoverTab[190048]"
	case elliptic.P384():
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:494
		_go_fuzz_dep_.CoverTab[190049]++
												return "P-384", nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:495
		// _ = "end of CoverTab[190049]"
	case elliptic.P521():
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:496
		_go_fuzz_dep_.CoverTab[190050]++
												return "P-521", nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:497
		// _ = "end of CoverTab[190050]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:498
		_go_fuzz_dep_.CoverTab[190051]++
												return "", fmt.Errorf("square/go-jose: unsupported/unknown elliptic curve")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:499
		// _ = "end of CoverTab[190051]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:500
	// _ = "end of CoverTab[190047]"
}

// Get size of curve in bytes
func curveSize(crv elliptic.Curve) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:504
	_go_fuzz_dep_.CoverTab[190052]++
											bits := crv.Params().BitSize

											div := bits / 8
											mod := bits % 8

											if mod == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:510
		_go_fuzz_dep_.CoverTab[190054]++
												return div
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:511
		// _ = "end of CoverTab[190054]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:512
		_go_fuzz_dep_.CoverTab[190055]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:512
		// _ = "end of CoverTab[190055]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:512
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:512
	// _ = "end of CoverTab[190052]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:512
	_go_fuzz_dep_.CoverTab[190053]++

											return div + 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:514
	// _ = "end of CoverTab[190053]"
}

func makeRawMessage(b []byte) *json.RawMessage {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:517
	_go_fuzz_dep_.CoverTab[190056]++
											rm := json.RawMessage(b)
											return &rm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:519
	// _ = "end of CoverTab[190056]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:520
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/shared.go:520
var _ = _go_fuzz_dep_.CoverTab
