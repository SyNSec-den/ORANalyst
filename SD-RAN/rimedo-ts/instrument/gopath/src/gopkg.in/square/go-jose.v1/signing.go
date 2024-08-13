//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:17
)

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"errors"
	"fmt"
)

// NonceSource represents a source of random nonces to go into JWS objects
type NonceSource interface {
	Nonce() (string, error)
}

// Signer represents a signer which takes a payload and produces a signed JWS object.
type Signer interface {
	Sign(payload []byte) (*JsonWebSignature, error)
	SetNonceSource(source NonceSource)
	SetEmbedJwk(embed bool)
}

// MultiSigner represents a signer which supports multiple recipients.
type MultiSigner interface {
	Sign(payload []byte) (*JsonWebSignature, error)
	SetNonceSource(source NonceSource)
	SetEmbedJwk(embed bool)
	AddRecipient(alg SignatureAlgorithm, signingKey interface{}) error
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
	embedJwk	bool
}

type recipientSigInfo struct {
	sigAlg		SignatureAlgorithm
	keyID		string
	publicKey	*JsonWebKey
	signer		payloadSigner
}

// NewSigner creates an appropriate signer based on the key type
func NewSigner(alg SignatureAlgorithm, signingKey interface{}) (Signer, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:68
	_go_fuzz_dep_.CoverTab[186583]++

											signer := NewMultiSigner()

											err := signer.AddRecipient(alg, signingKey)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:73
		_go_fuzz_dep_.CoverTab[186585]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:74
		// _ = "end of CoverTab[186585]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:75
		_go_fuzz_dep_.CoverTab[186586]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:75
		// _ = "end of CoverTab[186586]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:75
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:75
	// _ = "end of CoverTab[186583]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:75
	_go_fuzz_dep_.CoverTab[186584]++

											return signer, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:77
	// _ = "end of CoverTab[186584]"
}

// NewMultiSigner creates a signer for multiple recipients
func NewMultiSigner() MultiSigner {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:81
	_go_fuzz_dep_.CoverTab[186587]++
											return &genericSigner{
		recipients:	[]recipientSigInfo{},
		embedJwk:	true,
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:85
	// _ = "end of CoverTab[186587]"
}

// newVerifier creates a verifier based on the key type
func newVerifier(verificationKey interface{}) (payloadVerifier, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:89
	_go_fuzz_dep_.CoverTab[186588]++
											switch verificationKey := verificationKey.(type) {
	case *rsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:91
		_go_fuzz_dep_.CoverTab[186589]++
												return &rsaEncrypterVerifier{
			publicKey: verificationKey,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:94
		// _ = "end of CoverTab[186589]"
	case *ecdsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:95
		_go_fuzz_dep_.CoverTab[186590]++
												return &ecEncrypterVerifier{
			publicKey: verificationKey,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:98
		// _ = "end of CoverTab[186590]"
	case []byte:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:99
		_go_fuzz_dep_.CoverTab[186591]++
												return &symmetricMac{
			key: verificationKey,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:102
		// _ = "end of CoverTab[186591]"
	case *JsonWebKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:103
		_go_fuzz_dep_.CoverTab[186592]++
												return newVerifier(verificationKey.Key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:104
		// _ = "end of CoverTab[186592]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:105
		_go_fuzz_dep_.CoverTab[186593]++
												return nil, ErrUnsupportedKeyType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:106
		// _ = "end of CoverTab[186593]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:107
	// _ = "end of CoverTab[186588]"
}

func (ctx *genericSigner) AddRecipient(alg SignatureAlgorithm, signingKey interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:110
	_go_fuzz_dep_.CoverTab[186594]++
											recipient, err := makeJWSRecipient(alg, signingKey)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:112
		_go_fuzz_dep_.CoverTab[186596]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:113
		// _ = "end of CoverTab[186596]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:114
		_go_fuzz_dep_.CoverTab[186597]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:114
		// _ = "end of CoverTab[186597]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:114
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:114
	// _ = "end of CoverTab[186594]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:114
	_go_fuzz_dep_.CoverTab[186595]++

											ctx.recipients = append(ctx.recipients, recipient)
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:117
	// _ = "end of CoverTab[186595]"
}

func makeJWSRecipient(alg SignatureAlgorithm, signingKey interface{}) (recipientSigInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:120
	_go_fuzz_dep_.CoverTab[186598]++
											switch signingKey := signingKey.(type) {
	case *rsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:122
		_go_fuzz_dep_.CoverTab[186599]++
												return newRSASigner(alg, signingKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:123
		// _ = "end of CoverTab[186599]"
	case *ecdsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:124
		_go_fuzz_dep_.CoverTab[186600]++
												return newECDSASigner(alg, signingKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:125
		// _ = "end of CoverTab[186600]"
	case []byte:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:126
		_go_fuzz_dep_.CoverTab[186601]++
												return newSymmetricSigner(alg, signingKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:127
		// _ = "end of CoverTab[186601]"
	case *JsonWebKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:128
		_go_fuzz_dep_.CoverTab[186602]++
												recipient, err := makeJWSRecipient(alg, signingKey.Key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:130
			_go_fuzz_dep_.CoverTab[186605]++
													return recipientSigInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:131
			// _ = "end of CoverTab[186605]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:132
			_go_fuzz_dep_.CoverTab[186606]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:132
			// _ = "end of CoverTab[186606]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:132
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:132
		// _ = "end of CoverTab[186602]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:132
		_go_fuzz_dep_.CoverTab[186603]++
												recipient.keyID = signingKey.KeyID
												return recipient, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:134
		// _ = "end of CoverTab[186603]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:135
		_go_fuzz_dep_.CoverTab[186604]++
												return recipientSigInfo{}, ErrUnsupportedKeyType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:136
		// _ = "end of CoverTab[186604]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:137
	// _ = "end of CoverTab[186598]"
}

func (ctx *genericSigner) Sign(payload []byte) (*JsonWebSignature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:140
	_go_fuzz_dep_.CoverTab[186607]++
											obj := &JsonWebSignature{}
											obj.payload = payload
											obj.Signatures = make([]Signature, len(ctx.recipients))

											for i, recipient := range ctx.recipients {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:145
		_go_fuzz_dep_.CoverTab[186609]++
												protected := &rawHeader{
			Alg: string(recipient.sigAlg),
		}

		if recipient.publicKey != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:150
			_go_fuzz_dep_.CoverTab[186614]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:150
			return ctx.embedJwk
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:150
			// _ = "end of CoverTab[186614]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:150
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:150
			_go_fuzz_dep_.CoverTab[186615]++
													protected.Jwk = recipient.publicKey
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:151
			// _ = "end of CoverTab[186615]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:152
			_go_fuzz_dep_.CoverTab[186616]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:152
			// _ = "end of CoverTab[186616]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:152
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:152
		// _ = "end of CoverTab[186609]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:152
		_go_fuzz_dep_.CoverTab[186610]++
												if recipient.keyID != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:153
			_go_fuzz_dep_.CoverTab[186617]++
													protected.Kid = recipient.keyID
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:154
			// _ = "end of CoverTab[186617]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:155
			_go_fuzz_dep_.CoverTab[186618]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:155
			// _ = "end of CoverTab[186618]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:155
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:155
		// _ = "end of CoverTab[186610]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:155
		_go_fuzz_dep_.CoverTab[186611]++

												if ctx.nonceSource != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:157
			_go_fuzz_dep_.CoverTab[186619]++
													nonce, err := ctx.nonceSource.Nonce()
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:159
				_go_fuzz_dep_.CoverTab[186621]++
														return nil, fmt.Errorf("square/go-jose: Error generating nonce: %v", err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:160
				// _ = "end of CoverTab[186621]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:161
				_go_fuzz_dep_.CoverTab[186622]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:161
				// _ = "end of CoverTab[186622]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:161
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:161
			// _ = "end of CoverTab[186619]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:161
			_go_fuzz_dep_.CoverTab[186620]++
													protected.Nonce = nonce
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:162
			// _ = "end of CoverTab[186620]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:163
			_go_fuzz_dep_.CoverTab[186623]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:163
			// _ = "end of CoverTab[186623]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:163
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:163
		// _ = "end of CoverTab[186611]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:163
		_go_fuzz_dep_.CoverTab[186612]++

												serializedProtected := mustSerializeJSON(protected)

												input := []byte(fmt.Sprintf("%s.%s",
			base64URLEncode(serializedProtected),
			base64URLEncode(payload)))

		signatureInfo, err := recipient.signer.signPayload(input, recipient.sigAlg)
		if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:172
			_go_fuzz_dep_.CoverTab[186624]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:173
			// _ = "end of CoverTab[186624]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:174
			_go_fuzz_dep_.CoverTab[186625]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:174
			// _ = "end of CoverTab[186625]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:174
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:174
		// _ = "end of CoverTab[186612]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:174
		_go_fuzz_dep_.CoverTab[186613]++

												signatureInfo.protected = protected
												obj.Signatures[i] = signatureInfo
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:177
		// _ = "end of CoverTab[186613]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:178
	// _ = "end of CoverTab[186607]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:178
	_go_fuzz_dep_.CoverTab[186608]++

											return obj, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:180
	// _ = "end of CoverTab[186608]"
}

// SetNonceSource provides or updates a nonce pool to the first recipients.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:183
// After this method is called, the signer will consume one nonce per
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:183
// signature, returning an error it is unable to get a nonce.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:186
func (ctx *genericSigner) SetNonceSource(source NonceSource) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:186
	_go_fuzz_dep_.CoverTab[186626]++
											ctx.nonceSource = source
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:187
	// _ = "end of CoverTab[186626]"
}

// SetEmbedJwk specifies if the signing key should be embedded in the protected
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:190
// header, if any. It defaults to 'true', though that may change in the future.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:190
// Note that the use of embedded JWKs in the signature header can be dangerous,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:190
// as you cannot assume that the key received in a payload is trusted.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:194
func (ctx *genericSigner) SetEmbedJwk(embed bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:194
	_go_fuzz_dep_.CoverTab[186627]++
											ctx.embedJwk = embed
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:195
	// _ = "end of CoverTab[186627]"
}

// Verify validates the signature on the object and returns the payload.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:198
// This function does not support multi-signature, if you desire multi-sig
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:198
// verification use VerifyMulti instead.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:198
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:198
// Be careful when verifying signatures based on embedded JWKs inside the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:198
// payload header. You cannot assume that the key received in a payload is
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:198
// trusted.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:205
func (obj JsonWebSignature) Verify(verificationKey interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:205
	_go_fuzz_dep_.CoverTab[186628]++
											verifier, err := newVerifier(verificationKey)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:207
		_go_fuzz_dep_.CoverTab[186633]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:208
		// _ = "end of CoverTab[186633]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:209
		_go_fuzz_dep_.CoverTab[186634]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:209
		// _ = "end of CoverTab[186634]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:209
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:209
	// _ = "end of CoverTab[186628]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:209
	_go_fuzz_dep_.CoverTab[186629]++

											if len(obj.Signatures) > 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:211
		_go_fuzz_dep_.CoverTab[186635]++
												return nil, errors.New("square/go-jose: too many signatures in payload; expecting only one")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:212
		// _ = "end of CoverTab[186635]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:213
		_go_fuzz_dep_.CoverTab[186636]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:213
		// _ = "end of CoverTab[186636]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:213
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:213
	// _ = "end of CoverTab[186629]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:213
	_go_fuzz_dep_.CoverTab[186630]++

											signature := obj.Signatures[0]
											headers := signature.mergedHeaders()
											if len(headers.Crit) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:217
		_go_fuzz_dep_.CoverTab[186637]++

												return nil, ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:219
		// _ = "end of CoverTab[186637]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:220
		_go_fuzz_dep_.CoverTab[186638]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:220
		// _ = "end of CoverTab[186638]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:220
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:220
	// _ = "end of CoverTab[186630]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:220
	_go_fuzz_dep_.CoverTab[186631]++

											input := obj.computeAuthData(&signature)
											alg := SignatureAlgorithm(headers.Alg)
											err = verifier.verifyPayload(input, signature.Signature, alg)
											if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:225
		_go_fuzz_dep_.CoverTab[186639]++
												return obj.payload, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:226
		// _ = "end of CoverTab[186639]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:227
		_go_fuzz_dep_.CoverTab[186640]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:227
		// _ = "end of CoverTab[186640]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:227
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:227
	// _ = "end of CoverTab[186631]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:227
	_go_fuzz_dep_.CoverTab[186632]++

											return nil, ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:229
	// _ = "end of CoverTab[186632]"
}

// VerifyMulti validates (one of the multiple) signatures on the object and
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:232
// returns the index of the signature that was verified, along with the signature
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:232
// object and the payload. We return the signature and index to guarantee that
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:232
// callers are getting the verified value.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:236
func (obj JsonWebSignature) VerifyMulti(verificationKey interface{}) (int, Signature, []byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:236
	_go_fuzz_dep_.CoverTab[186641]++
											verifier, err := newVerifier(verificationKey)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:238
		_go_fuzz_dep_.CoverTab[186644]++
												return -1, Signature{}, nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:239
		// _ = "end of CoverTab[186644]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:240
		_go_fuzz_dep_.CoverTab[186645]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:240
		// _ = "end of CoverTab[186645]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:240
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:240
	// _ = "end of CoverTab[186641]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:240
	_go_fuzz_dep_.CoverTab[186642]++

											for i, signature := range obj.Signatures {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:242
		_go_fuzz_dep_.CoverTab[186646]++
												headers := signature.mergedHeaders()
												if len(headers.Crit) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:244
			_go_fuzz_dep_.CoverTab[186648]++

													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:246
			// _ = "end of CoverTab[186648]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:247
			_go_fuzz_dep_.CoverTab[186649]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:247
			// _ = "end of CoverTab[186649]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:247
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:247
		// _ = "end of CoverTab[186646]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:247
		_go_fuzz_dep_.CoverTab[186647]++

												input := obj.computeAuthData(&signature)
												alg := SignatureAlgorithm(headers.Alg)
												err := verifier.verifyPayload(input, signature.Signature, alg)
												if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:252
			_go_fuzz_dep_.CoverTab[186650]++
													return i, signature, obj.payload, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:253
			// _ = "end of CoverTab[186650]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:254
			_go_fuzz_dep_.CoverTab[186651]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:254
			// _ = "end of CoverTab[186651]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:254
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:254
		// _ = "end of CoverTab[186647]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:255
	// _ = "end of CoverTab[186642]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:255
	_go_fuzz_dep_.CoverTab[186643]++

											return -1, Signature{}, nil, ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:257
	// _ = "end of CoverTab[186643]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:258
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/signing.go:258
var _ = _go_fuzz_dep_.CoverTab
