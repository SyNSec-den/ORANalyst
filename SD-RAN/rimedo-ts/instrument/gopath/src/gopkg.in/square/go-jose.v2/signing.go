//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:17
)

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"

	"golang.org/x/crypto/ed25519"

	"gopkg.in/square/go-jose.v2/json"
)

// NonceSource represents a source of random nonces to go into JWS objects
type NonceSource interface {
	Nonce() (string, error)
}

// Signer represents a signer which takes a payload and produces a signed JWS object.
type Signer interface {
	Sign(payload []byte) (*JSONWebSignature, error)
	Options() SignerOptions
}

// SigningKey represents an algorithm/key used to sign a message.
type SigningKey struct {
	Algorithm	SignatureAlgorithm
	Key		interface{}
}

// SignerOptions represents options that can be set when creating signers.
type SignerOptions struct {
	NonceSource	NonceSource
	EmbedJWK	bool

	// Optional map of additional keys to be inserted into the protected header
	// of a JWS object. Some specifications which make use of JWS like to insert
	// additional values here. All values must be JSON-serializable.
	ExtraHeaders	map[HeaderKey]interface{}
}

// WithHeader adds an arbitrary value to the ExtraHeaders map, initializing it
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:60
// if necessary. It returns itself and so can be used in a fluent style.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:62
func (so *SignerOptions) WithHeader(k HeaderKey, v interface{}) *SignerOptions {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:62
	_go_fuzz_dep_.CoverTab[190057]++
											if so.ExtraHeaders == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:63
		_go_fuzz_dep_.CoverTab[190059]++
												so.ExtraHeaders = map[HeaderKey]interface{}{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:64
		// _ = "end of CoverTab[190059]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:65
		_go_fuzz_dep_.CoverTab[190060]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:65
		// _ = "end of CoverTab[190060]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:65
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:65
	// _ = "end of CoverTab[190057]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:65
	_go_fuzz_dep_.CoverTab[190058]++
											so.ExtraHeaders[k] = v
											return so
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:67
	// _ = "end of CoverTab[190058]"
}

// WithContentType adds a content type ("cty") header and returns the updated
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:70
// SignerOptions.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:72
func (so *SignerOptions) WithContentType(contentType ContentType) *SignerOptions {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:72
	_go_fuzz_dep_.CoverTab[190061]++
											return so.WithHeader(HeaderContentType, contentType)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:73
	// _ = "end of CoverTab[190061]"
}

// WithType adds a type ("typ") header and returns the updated SignerOptions.
func (so *SignerOptions) WithType(typ ContentType) *SignerOptions {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:77
	_go_fuzz_dep_.CoverTab[190062]++
											return so.WithHeader(HeaderType, typ)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:78
	// _ = "end of CoverTab[190062]"
}

// WithCritical adds the given names to the critical ("crit") header and returns
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:81
// the updated SignerOptions.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:83
func (so *SignerOptions) WithCritical(names ...string) *SignerOptions {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:83
	_go_fuzz_dep_.CoverTab[190063]++
											if so.ExtraHeaders[headerCritical] == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:84
		_go_fuzz_dep_.CoverTab[190065]++
												so.WithHeader(headerCritical, make([]string, 0, len(names)))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:85
		// _ = "end of CoverTab[190065]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:86
		_go_fuzz_dep_.CoverTab[190066]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:86
		// _ = "end of CoverTab[190066]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:86
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:86
	// _ = "end of CoverTab[190063]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:86
	_go_fuzz_dep_.CoverTab[190064]++
											crit := so.ExtraHeaders[headerCritical].([]string)
											so.ExtraHeaders[headerCritical] = append(crit, names...)
											return so
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:89
	// _ = "end of CoverTab[190064]"
}

// WithBase64 adds a base64url-encode payload ("b64") header and returns the updated
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:92
// SignerOptions. When the "b64" value is "false", the payload is not base64 encoded.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:94
func (so *SignerOptions) WithBase64(b64 bool) *SignerOptions {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:94
	_go_fuzz_dep_.CoverTab[190067]++
											if !b64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:95
		_go_fuzz_dep_.CoverTab[190069]++
												so.WithHeader(headerB64, b64)
												so.WithCritical(headerB64)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:97
		// _ = "end of CoverTab[190069]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:98
		_go_fuzz_dep_.CoverTab[190070]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:98
		// _ = "end of CoverTab[190070]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:98
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:98
	// _ = "end of CoverTab[190067]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:98
	_go_fuzz_dep_.CoverTab[190068]++
											return so
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:99
	// _ = "end of CoverTab[190068]"
}

type payloadSigner interface {
	signPayload(payload []byte, alg SignatureAlgorithm) (Signature, error)
}

type payloadVerifier interface {
	verifyPayload(payload []byte, signature []byte, alg SignatureAlgorithm) error
}

type genericSigner struct {
	recipients	[]recipientSigInfo
	nonceSource	NonceSource
	embedJWK	bool
	extraHeaders	map[HeaderKey]interface{}
}

type recipientSigInfo struct {
	sigAlg		SignatureAlgorithm
	publicKey	func() *JSONWebKey
	signer		payloadSigner
}

func staticPublicKey(jwk *JSONWebKey) func() *JSONWebKey {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:123
	_go_fuzz_dep_.CoverTab[190071]++
											return func() *JSONWebKey {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:124
		_go_fuzz_dep_.CoverTab[190072]++
												return jwk
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:125
		// _ = "end of CoverTab[190072]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:126
	// _ = "end of CoverTab[190071]"
}

// NewSigner creates an appropriate signer based on the key type
func NewSigner(sig SigningKey, opts *SignerOptions) (Signer, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:130
	_go_fuzz_dep_.CoverTab[190073]++
											return NewMultiSigner([]SigningKey{sig}, opts)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:131
	// _ = "end of CoverTab[190073]"
}

// NewMultiSigner creates a signer for multiple recipients
func NewMultiSigner(sigs []SigningKey, opts *SignerOptions) (Signer, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:135
	_go_fuzz_dep_.CoverTab[190074]++
											signer := &genericSigner{recipients: []recipientSigInfo{}}

											if opts != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:138
		_go_fuzz_dep_.CoverTab[190077]++
												signer.nonceSource = opts.NonceSource
												signer.embedJWK = opts.EmbedJWK
												signer.extraHeaders = opts.ExtraHeaders
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:141
		// _ = "end of CoverTab[190077]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:142
		_go_fuzz_dep_.CoverTab[190078]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:142
		// _ = "end of CoverTab[190078]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:142
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:142
	// _ = "end of CoverTab[190074]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:142
	_go_fuzz_dep_.CoverTab[190075]++

											for _, sig := range sigs {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:144
		_go_fuzz_dep_.CoverTab[190079]++
												err := signer.addRecipient(sig.Algorithm, sig.Key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:146
			_go_fuzz_dep_.CoverTab[190080]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:147
			// _ = "end of CoverTab[190080]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:148
			_go_fuzz_dep_.CoverTab[190081]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:148
			// _ = "end of CoverTab[190081]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:148
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:148
		// _ = "end of CoverTab[190079]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:149
	// _ = "end of CoverTab[190075]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:149
	_go_fuzz_dep_.CoverTab[190076]++

											return signer, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:151
	// _ = "end of CoverTab[190076]"
}

// newVerifier creates a verifier based on the key type
func newVerifier(verificationKey interface{}) (payloadVerifier, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:155
	_go_fuzz_dep_.CoverTab[190082]++
											switch verificationKey := verificationKey.(type) {
	case ed25519.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:157
		_go_fuzz_dep_.CoverTab[190085]++
												return &edEncrypterVerifier{
			publicKey: verificationKey,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:160
		// _ = "end of CoverTab[190085]"
	case *rsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:161
		_go_fuzz_dep_.CoverTab[190086]++
												return &rsaEncrypterVerifier{
			publicKey: verificationKey,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:164
		// _ = "end of CoverTab[190086]"
	case *ecdsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:165
		_go_fuzz_dep_.CoverTab[190087]++
												return &ecEncrypterVerifier{
			publicKey: verificationKey,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:168
		// _ = "end of CoverTab[190087]"
	case []byte:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:169
		_go_fuzz_dep_.CoverTab[190088]++
												return &symmetricMac{
			key: verificationKey,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:172
		// _ = "end of CoverTab[190088]"
	case JSONWebKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:173
		_go_fuzz_dep_.CoverTab[190089]++
												return newVerifier(verificationKey.Key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:174
		// _ = "end of CoverTab[190089]"
	case *JSONWebKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:175
		_go_fuzz_dep_.CoverTab[190090]++
												return newVerifier(verificationKey.Key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:176
		// _ = "end of CoverTab[190090]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:177
	// _ = "end of CoverTab[190082]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:177
	_go_fuzz_dep_.CoverTab[190083]++
											if ov, ok := verificationKey.(OpaqueVerifier); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:178
		_go_fuzz_dep_.CoverTab[190091]++
												return &opaqueVerifier{verifier: ov}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:179
		// _ = "end of CoverTab[190091]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:180
		_go_fuzz_dep_.CoverTab[190092]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:180
		// _ = "end of CoverTab[190092]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:180
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:180
	// _ = "end of CoverTab[190083]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:180
	_go_fuzz_dep_.CoverTab[190084]++
											return nil, ErrUnsupportedKeyType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:181
	// _ = "end of CoverTab[190084]"
}

func (ctx *genericSigner) addRecipient(alg SignatureAlgorithm, signingKey interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:184
	_go_fuzz_dep_.CoverTab[190093]++
											recipient, err := makeJWSRecipient(alg, signingKey)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:186
		_go_fuzz_dep_.CoverTab[190095]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:187
		// _ = "end of CoverTab[190095]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:188
		_go_fuzz_dep_.CoverTab[190096]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:188
		// _ = "end of CoverTab[190096]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:188
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:188
	// _ = "end of CoverTab[190093]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:188
	_go_fuzz_dep_.CoverTab[190094]++

											ctx.recipients = append(ctx.recipients, recipient)
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:191
	// _ = "end of CoverTab[190094]"
}

func makeJWSRecipient(alg SignatureAlgorithm, signingKey interface{}) (recipientSigInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:194
	_go_fuzz_dep_.CoverTab[190097]++
											switch signingKey := signingKey.(type) {
	case ed25519.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:196
		_go_fuzz_dep_.CoverTab[190100]++
												return newEd25519Signer(alg, signingKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:197
		// _ = "end of CoverTab[190100]"
	case *rsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:198
		_go_fuzz_dep_.CoverTab[190101]++
												return newRSASigner(alg, signingKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:199
		// _ = "end of CoverTab[190101]"
	case *ecdsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:200
		_go_fuzz_dep_.CoverTab[190102]++
												return newECDSASigner(alg, signingKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:201
		// _ = "end of CoverTab[190102]"
	case []byte:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:202
		_go_fuzz_dep_.CoverTab[190103]++
												return newSymmetricSigner(alg, signingKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:203
		// _ = "end of CoverTab[190103]"
	case JSONWebKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:204
		_go_fuzz_dep_.CoverTab[190104]++
												return newJWKSigner(alg, signingKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:205
		// _ = "end of CoverTab[190104]"
	case *JSONWebKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:206
		_go_fuzz_dep_.CoverTab[190105]++
												return newJWKSigner(alg, *signingKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:207
		// _ = "end of CoverTab[190105]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:208
	// _ = "end of CoverTab[190097]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:208
	_go_fuzz_dep_.CoverTab[190098]++
											if signer, ok := signingKey.(OpaqueSigner); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:209
		_go_fuzz_dep_.CoverTab[190106]++
												return newOpaqueSigner(alg, signer)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:210
		// _ = "end of CoverTab[190106]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:211
		_go_fuzz_dep_.CoverTab[190107]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:211
		// _ = "end of CoverTab[190107]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:211
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:211
	// _ = "end of CoverTab[190098]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:211
	_go_fuzz_dep_.CoverTab[190099]++
											return recipientSigInfo{}, ErrUnsupportedKeyType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:212
	// _ = "end of CoverTab[190099]"
}

func newJWKSigner(alg SignatureAlgorithm, signingKey JSONWebKey) (recipientSigInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:215
	_go_fuzz_dep_.CoverTab[190108]++
											recipient, err := makeJWSRecipient(alg, signingKey.Key)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:217
		_go_fuzz_dep_.CoverTab[190111]++
												return recipientSigInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:218
		// _ = "end of CoverTab[190111]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:219
		_go_fuzz_dep_.CoverTab[190112]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:219
		// _ = "end of CoverTab[190112]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:219
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:219
	// _ = "end of CoverTab[190108]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:219
	_go_fuzz_dep_.CoverTab[190109]++
											if recipient.publicKey != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:220
		_go_fuzz_dep_.CoverTab[190113]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:220
		return recipient.publicKey() != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:220
		// _ = "end of CoverTab[190113]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:220
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:220
		_go_fuzz_dep_.CoverTab[190114]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:224
		publicKey := signingKey
												publicKey.Key = recipient.publicKey().Key
												recipient.publicKey = staticPublicKey(&publicKey)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:229
		if !recipient.publicKey().IsPublic() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:229
			_go_fuzz_dep_.CoverTab[190115]++
													return recipientSigInfo{}, errors.New("square/go-jose: public key was unexpectedly not public")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:230
			// _ = "end of CoverTab[190115]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:231
			_go_fuzz_dep_.CoverTab[190116]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:231
			// _ = "end of CoverTab[190116]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:231
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:231
		// _ = "end of CoverTab[190114]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:232
		_go_fuzz_dep_.CoverTab[190117]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:232
		// _ = "end of CoverTab[190117]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:232
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:232
	// _ = "end of CoverTab[190109]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:232
	_go_fuzz_dep_.CoverTab[190110]++
											return recipient, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:233
	// _ = "end of CoverTab[190110]"
}

func (ctx *genericSigner) Sign(payload []byte) (*JSONWebSignature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:236
	_go_fuzz_dep_.CoverTab[190118]++
											obj := &JSONWebSignature{}
											obj.payload = payload
											obj.Signatures = make([]Signature, len(ctx.recipients))

											for i, recipient := range ctx.recipients {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:241
		_go_fuzz_dep_.CoverTab[190120]++
												protected := map[HeaderKey]interface{}{
			headerAlgorithm: string(recipient.sigAlg),
		}

		if recipient.publicKey != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:246
			_go_fuzz_dep_.CoverTab[190128]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:246
			return recipient.publicKey() != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:246
			// _ = "end of CoverTab[190128]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:246
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:246
			_go_fuzz_dep_.CoverTab[190129]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:255
			if ctx.embedJWK {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:255
				_go_fuzz_dep_.CoverTab[190130]++
														protected[headerJWK] = recipient.publicKey()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:256
				// _ = "end of CoverTab[190130]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:257
				_go_fuzz_dep_.CoverTab[190131]++
														keyID := recipient.publicKey().KeyID
														if keyID != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:259
					_go_fuzz_dep_.CoverTab[190132]++
															protected[headerKeyID] = keyID
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:260
					// _ = "end of CoverTab[190132]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:261
					_go_fuzz_dep_.CoverTab[190133]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:261
					// _ = "end of CoverTab[190133]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:261
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:261
				// _ = "end of CoverTab[190131]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:262
			// _ = "end of CoverTab[190129]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:263
			_go_fuzz_dep_.CoverTab[190134]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:263
			// _ = "end of CoverTab[190134]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:263
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:263
		// _ = "end of CoverTab[190120]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:263
		_go_fuzz_dep_.CoverTab[190121]++

												if ctx.nonceSource != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:265
			_go_fuzz_dep_.CoverTab[190135]++
													nonce, err := ctx.nonceSource.Nonce()
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:267
				_go_fuzz_dep_.CoverTab[190137]++
														return nil, fmt.Errorf("square/go-jose: Error generating nonce: %v", err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:268
				// _ = "end of CoverTab[190137]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:269
				_go_fuzz_dep_.CoverTab[190138]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:269
				// _ = "end of CoverTab[190138]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:269
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:269
			// _ = "end of CoverTab[190135]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:269
			_go_fuzz_dep_.CoverTab[190136]++
													protected[headerNonce] = nonce
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:270
			// _ = "end of CoverTab[190136]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:271
			_go_fuzz_dep_.CoverTab[190139]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:271
			// _ = "end of CoverTab[190139]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:271
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:271
		// _ = "end of CoverTab[190121]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:271
		_go_fuzz_dep_.CoverTab[190122]++

												for k, v := range ctx.extraHeaders {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:273
			_go_fuzz_dep_.CoverTab[190140]++
													protected[k] = v
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:274
			// _ = "end of CoverTab[190140]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:275
		// _ = "end of CoverTab[190122]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:275
		_go_fuzz_dep_.CoverTab[190123]++

												serializedProtected := mustSerializeJSON(protected)
												needsBase64 := true

												if b64, ok := protected[headerB64]; ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:280
			_go_fuzz_dep_.CoverTab[190141]++
													if needsBase64, ok = b64.(bool); !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:281
				_go_fuzz_dep_.CoverTab[190142]++
														return nil, errors.New("square/go-jose: Invalid b64 header parameter")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:282
				// _ = "end of CoverTab[190142]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:283
				_go_fuzz_dep_.CoverTab[190143]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:283
				// _ = "end of CoverTab[190143]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:283
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:283
			// _ = "end of CoverTab[190141]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:284
			_go_fuzz_dep_.CoverTab[190144]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:284
			// _ = "end of CoverTab[190144]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:284
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:284
		// _ = "end of CoverTab[190123]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:284
		_go_fuzz_dep_.CoverTab[190124]++

												var input bytes.Buffer

												input.WriteString(base64.RawURLEncoding.EncodeToString(serializedProtected))
												input.WriteByte('.')

												if needsBase64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:291
			_go_fuzz_dep_.CoverTab[190145]++
													input.WriteString(base64.RawURLEncoding.EncodeToString(payload))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:292
			// _ = "end of CoverTab[190145]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:293
			_go_fuzz_dep_.CoverTab[190146]++
													input.Write(payload)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:294
			// _ = "end of CoverTab[190146]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:295
		// _ = "end of CoverTab[190124]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:295
		_go_fuzz_dep_.CoverTab[190125]++

												signatureInfo, err := recipient.signer.signPayload(input.Bytes(), recipient.sigAlg)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:298
			_go_fuzz_dep_.CoverTab[190147]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:299
			// _ = "end of CoverTab[190147]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:300
			_go_fuzz_dep_.CoverTab[190148]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:300
			// _ = "end of CoverTab[190148]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:300
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:300
		// _ = "end of CoverTab[190125]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:300
		_go_fuzz_dep_.CoverTab[190126]++

												signatureInfo.protected = &rawHeader{}
												for k, v := range protected {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:303
			_go_fuzz_dep_.CoverTab[190149]++
													b, err := json.Marshal(v)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:305
				_go_fuzz_dep_.CoverTab[190151]++
														return nil, fmt.Errorf("square/go-jose: Error marshalling item %#v: %v", k, err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:306
				// _ = "end of CoverTab[190151]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:307
				_go_fuzz_dep_.CoverTab[190152]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:307
				// _ = "end of CoverTab[190152]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:307
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:307
			// _ = "end of CoverTab[190149]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:307
			_go_fuzz_dep_.CoverTab[190150]++
													(*signatureInfo.protected)[k] = makeRawMessage(b)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:308
			// _ = "end of CoverTab[190150]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:309
		// _ = "end of CoverTab[190126]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:309
		_go_fuzz_dep_.CoverTab[190127]++
												obj.Signatures[i] = signatureInfo
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:310
		// _ = "end of CoverTab[190127]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:311
	// _ = "end of CoverTab[190118]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:311
	_go_fuzz_dep_.CoverTab[190119]++

											return obj, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:313
	// _ = "end of CoverTab[190119]"
}

func (ctx *genericSigner) Options() SignerOptions {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:316
	_go_fuzz_dep_.CoverTab[190153]++
											return SignerOptions{
		NonceSource:	ctx.nonceSource,
		EmbedJWK:	ctx.embedJWK,
		ExtraHeaders:	ctx.extraHeaders,
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:321
	// _ = "end of CoverTab[190153]"
}

// Verify validates the signature on the object and returns the payload.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:324
// This function does not support multi-signature, if you desire multi-sig
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:324
// verification use VerifyMulti instead.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:324
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:324
// Be careful when verifying signatures based on embedded JWKs inside the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:324
// payload header. You cannot assume that the key received in a payload is
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:324
// trusted.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:331
func (obj JSONWebSignature) Verify(verificationKey interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:331
	_go_fuzz_dep_.CoverTab[190154]++
											err := obj.DetachedVerify(obj.payload, verificationKey)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:333
		_go_fuzz_dep_.CoverTab[190156]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:334
		// _ = "end of CoverTab[190156]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:335
		_go_fuzz_dep_.CoverTab[190157]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:335
		// _ = "end of CoverTab[190157]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:335
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:335
	// _ = "end of CoverTab[190154]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:335
	_go_fuzz_dep_.CoverTab[190155]++
											return obj.payload, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:336
	// _ = "end of CoverTab[190155]"
}

// UnsafePayloadWithoutVerification returns the payload without
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:339
// verifying it. The content returned from this function cannot be
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:339
// trusted.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:342
func (obj JSONWebSignature) UnsafePayloadWithoutVerification() []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:342
	_go_fuzz_dep_.CoverTab[190158]++
											return obj.payload
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:343
	// _ = "end of CoverTab[190158]"
}

// DetachedVerify validates a detached signature on the given payload. In
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:346
// most cases, you will probably want to use Verify instead. DetachedVerify
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:346
// is only useful if you have a payload and signature that are separated from
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:346
// each other.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:350
func (obj JSONWebSignature) DetachedVerify(payload []byte, verificationKey interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:350
	_go_fuzz_dep_.CoverTab[190159]++
											verifier, err := newVerifier(verificationKey)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:352
		_go_fuzz_dep_.CoverTab[190166]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:353
		// _ = "end of CoverTab[190166]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:354
		_go_fuzz_dep_.CoverTab[190167]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:354
		// _ = "end of CoverTab[190167]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:354
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:354
	// _ = "end of CoverTab[190159]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:354
	_go_fuzz_dep_.CoverTab[190160]++

											if len(obj.Signatures) > 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:356
		_go_fuzz_dep_.CoverTab[190168]++
												return errors.New("square/go-jose: too many signatures in payload; expecting only one")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:357
		// _ = "end of CoverTab[190168]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:358
		_go_fuzz_dep_.CoverTab[190169]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:358
		// _ = "end of CoverTab[190169]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:358
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:358
	// _ = "end of CoverTab[190160]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:358
	_go_fuzz_dep_.CoverTab[190161]++

											signature := obj.Signatures[0]
											headers := signature.mergedHeaders()
											critical, err := headers.getCritical()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:363
		_go_fuzz_dep_.CoverTab[190170]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:364
		// _ = "end of CoverTab[190170]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:365
		_go_fuzz_dep_.CoverTab[190171]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:365
		// _ = "end of CoverTab[190171]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:365
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:365
	// _ = "end of CoverTab[190161]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:365
	_go_fuzz_dep_.CoverTab[190162]++

											for _, name := range critical {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:367
		_go_fuzz_dep_.CoverTab[190172]++
												if !supportedCritical[name] {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:368
			_go_fuzz_dep_.CoverTab[190173]++
													return ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:369
			// _ = "end of CoverTab[190173]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:370
			_go_fuzz_dep_.CoverTab[190174]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:370
			// _ = "end of CoverTab[190174]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:370
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:370
		// _ = "end of CoverTab[190172]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:371
	// _ = "end of CoverTab[190162]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:371
	_go_fuzz_dep_.CoverTab[190163]++

											input, err := obj.computeAuthData(payload, &signature)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:374
		_go_fuzz_dep_.CoverTab[190175]++
												return ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:375
		// _ = "end of CoverTab[190175]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:376
		_go_fuzz_dep_.CoverTab[190176]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:376
		// _ = "end of CoverTab[190176]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:376
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:376
	// _ = "end of CoverTab[190163]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:376
	_go_fuzz_dep_.CoverTab[190164]++

											alg := headers.getSignatureAlgorithm()
											err = verifier.verifyPayload(input, signature.Signature, alg)
											if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:380
		_go_fuzz_dep_.CoverTab[190177]++
												return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:381
		// _ = "end of CoverTab[190177]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:382
		_go_fuzz_dep_.CoverTab[190178]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:382
		// _ = "end of CoverTab[190178]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:382
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:382
	// _ = "end of CoverTab[190164]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:382
	_go_fuzz_dep_.CoverTab[190165]++

											return ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:384
	// _ = "end of CoverTab[190165]"
}

// VerifyMulti validates (one of the multiple) signatures on the object and
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:387
// returns the index of the signature that was verified, along with the signature
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:387
// object and the payload. We return the signature and index to guarantee that
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:387
// callers are getting the verified value.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:391
func (obj JSONWebSignature) VerifyMulti(verificationKey interface{}) (int, Signature, []byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:391
	_go_fuzz_dep_.CoverTab[190179]++
											idx, sig, err := obj.DetachedVerifyMulti(obj.payload, verificationKey)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:393
		_go_fuzz_dep_.CoverTab[190181]++
												return -1, Signature{}, nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:394
		// _ = "end of CoverTab[190181]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:395
		_go_fuzz_dep_.CoverTab[190182]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:395
		// _ = "end of CoverTab[190182]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:395
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:395
	// _ = "end of CoverTab[190179]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:395
	_go_fuzz_dep_.CoverTab[190180]++
											return idx, sig, obj.payload, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:396
	// _ = "end of CoverTab[190180]"
}

// DetachedVerifyMulti validates a detached signature on the given payload with
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:399
// a signature/object that has potentially multiple signers. This returns the index
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:399
// of the signature that was verified, along with the signature object. We return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:399
// the signature and index to guarantee that callers are getting the verified value.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:399
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:399
// In most cases, you will probably want to use Verify or VerifyMulti instead.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:399
// DetachedVerifyMulti is only useful if you have a payload and signature that are
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:399
// separated from each other, and the signature can have multiple signers at the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:399
// same time.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:408
func (obj JSONWebSignature) DetachedVerifyMulti(payload []byte, verificationKey interface{}) (int, Signature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:408
	_go_fuzz_dep_.CoverTab[190183]++
											verifier, err := newVerifier(verificationKey)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:410
		_go_fuzz_dep_.CoverTab[190186]++
												return -1, Signature{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:411
		// _ = "end of CoverTab[190186]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:412
		_go_fuzz_dep_.CoverTab[190187]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:412
		// _ = "end of CoverTab[190187]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:412
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:412
	// _ = "end of CoverTab[190183]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:412
	_go_fuzz_dep_.CoverTab[190184]++

outer:
	for i, signature := range obj.Signatures {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:415
		_go_fuzz_dep_.CoverTab[190188]++
												headers := signature.mergedHeaders()
												critical, err := headers.getCritical()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:418
			_go_fuzz_dep_.CoverTab[190192]++
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:419
			// _ = "end of CoverTab[190192]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:420
			_go_fuzz_dep_.CoverTab[190193]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:420
			// _ = "end of CoverTab[190193]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:420
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:420
		// _ = "end of CoverTab[190188]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:420
		_go_fuzz_dep_.CoverTab[190189]++

												for _, name := range critical {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:422
			_go_fuzz_dep_.CoverTab[190194]++
													if !supportedCritical[name] {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:423
				_go_fuzz_dep_.CoverTab[190195]++
														continue outer
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:424
				// _ = "end of CoverTab[190195]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:425
				_go_fuzz_dep_.CoverTab[190196]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:425
				// _ = "end of CoverTab[190196]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:425
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:425
			// _ = "end of CoverTab[190194]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:426
		// _ = "end of CoverTab[190189]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:426
		_go_fuzz_dep_.CoverTab[190190]++

												input, err := obj.computeAuthData(payload, &signature)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:429
			_go_fuzz_dep_.CoverTab[190197]++
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:430
			// _ = "end of CoverTab[190197]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:431
			_go_fuzz_dep_.CoverTab[190198]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:431
			// _ = "end of CoverTab[190198]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:431
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:431
		// _ = "end of CoverTab[190190]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:431
		_go_fuzz_dep_.CoverTab[190191]++

												alg := headers.getSignatureAlgorithm()
												err = verifier.verifyPayload(input, signature.Signature, alg)
												if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:435
			_go_fuzz_dep_.CoverTab[190199]++
													return i, signature, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:436
			// _ = "end of CoverTab[190199]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:437
			_go_fuzz_dep_.CoverTab[190200]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:437
			// _ = "end of CoverTab[190200]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:437
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:437
		// _ = "end of CoverTab[190191]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:438
	// _ = "end of CoverTab[190184]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:438
	_go_fuzz_dep_.CoverTab[190185]++

											return -1, Signature{}, ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:440
	// _ = "end of CoverTab[190185]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:441
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/signing.go:441
var _ = _go_fuzz_dep_.CoverTab
