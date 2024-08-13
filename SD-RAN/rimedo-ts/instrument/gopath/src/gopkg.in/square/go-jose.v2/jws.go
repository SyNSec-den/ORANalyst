//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:17
)

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"gopkg.in/square/go-jose.v2/json"
)

// rawJSONWebSignature represents a raw JWS JSON object. Used for parsing/serializing.
type rawJSONWebSignature struct {
	Payload		*byteBuffer		`json:"payload,omitempty"`
	Signatures	[]rawSignatureInfo	`json:"signatures,omitempty"`
	Protected	*byteBuffer		`json:"protected,omitempty"`
	Header		*rawHeader		`json:"header,omitempty"`
	Signature	*byteBuffer		`json:"signature,omitempty"`
}

// rawSignatureInfo represents a single JWS signature over the JWS payload and protected header.
type rawSignatureInfo struct {
	Protected	*byteBuffer	`json:"protected,omitempty"`
	Header		*rawHeader	`json:"header,omitempty"`
	Signature	*byteBuffer	`json:"signature,omitempty"`
}

// JSONWebSignature represents a signed JWS object after parsing.
type JSONWebSignature struct {
	payload	[]byte
	// Signatures attached to this object (may be more than one for multi-sig).
	// Be careful about accessing these directly, prefer to use Verify() or
	// VerifyMulti() to ensure that the data you're getting is verified.
	Signatures	[]Signature
}

// Signature represents a single signature over the JWS payload and protected header.
type Signature struct {
	// Merged header fields. Contains both protected and unprotected header
	// values. Prefer using Protected and Unprotected fields instead of this.
	// Values in this header may or may not have been signed and in general
	// should not be trusted.
	Header	Header

	// Protected header. Values in this header were signed and
	// will be verified as part of the signature verification process.
	Protected	Header

	// Unprotected header. Values in this header were not signed
	// and in general should not be trusted.
	Unprotected	Header

	// The actual signature value
	Signature	[]byte

	protected	*rawHeader
	header		*rawHeader
	original	*rawSignatureInfo
}

// ParseSigned parses a signed message in compact or full serialization format.
func ParseSigned(signature string) (*JSONWebSignature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:79
	_go_fuzz_dep_.CoverTab[189752]++
										signature = stripWhitespace(signature)
										if strings.HasPrefix(signature, "{") {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:81
		_go_fuzz_dep_.CoverTab[189754]++
											return parseSignedFull(signature)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:82
		// _ = "end of CoverTab[189754]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:83
		_go_fuzz_dep_.CoverTab[189755]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:83
		// _ = "end of CoverTab[189755]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:83
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:83
	// _ = "end of CoverTab[189752]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:83
	_go_fuzz_dep_.CoverTab[189753]++

										return parseSignedCompact(signature, nil)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:85
	// _ = "end of CoverTab[189753]"
}

// ParseDetached parses a signed message in compact serialization format with detached payload.
func ParseDetached(signature string, payload []byte) (*JSONWebSignature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:89
	_go_fuzz_dep_.CoverTab[189756]++
										if payload == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:90
		_go_fuzz_dep_.CoverTab[189758]++
											return nil, errors.New("square/go-jose: nil payload")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:91
		// _ = "end of CoverTab[189758]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:92
		_go_fuzz_dep_.CoverTab[189759]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:92
		// _ = "end of CoverTab[189759]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:92
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:92
	// _ = "end of CoverTab[189756]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:92
	_go_fuzz_dep_.CoverTab[189757]++
										return parseSignedCompact(stripWhitespace(signature), payload)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:93
	// _ = "end of CoverTab[189757]"
}

// Get a header value
func (sig Signature) mergedHeaders() rawHeader {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:97
	_go_fuzz_dep_.CoverTab[189760]++
										out := rawHeader{}
										out.merge(sig.protected)
										out.merge(sig.header)
										return out
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:101
	// _ = "end of CoverTab[189760]"
}

// Compute data to be signed
func (obj JSONWebSignature) computeAuthData(payload []byte, signature *Signature) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:105
	_go_fuzz_dep_.CoverTab[189761]++
										var authData bytes.Buffer

										protectedHeader := new(rawHeader)

										if signature.original != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:110
		_go_fuzz_dep_.CoverTab[189765]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:110
		return signature.original.Protected != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:110
		// _ = "end of CoverTab[189765]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:110
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:110
		_go_fuzz_dep_.CoverTab[189766]++
											if err := json.Unmarshal(signature.original.Protected.bytes(), protectedHeader); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:111
			_go_fuzz_dep_.CoverTab[189768]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:112
			// _ = "end of CoverTab[189768]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:113
			_go_fuzz_dep_.CoverTab[189769]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:113
			// _ = "end of CoverTab[189769]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:113
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:113
		// _ = "end of CoverTab[189766]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:113
		_go_fuzz_dep_.CoverTab[189767]++
											authData.WriteString(signature.original.Protected.base64())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:114
		// _ = "end of CoverTab[189767]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:115
		_go_fuzz_dep_.CoverTab[189770]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:115
		if signature.protected != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:115
			_go_fuzz_dep_.CoverTab[189771]++
												protectedHeader = signature.protected
												authData.WriteString(base64.RawURLEncoding.EncodeToString(mustSerializeJSON(protectedHeader)))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:117
			// _ = "end of CoverTab[189771]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:118
			_go_fuzz_dep_.CoverTab[189772]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:118
			// _ = "end of CoverTab[189772]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:118
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:118
		// _ = "end of CoverTab[189770]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:118
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:118
	// _ = "end of CoverTab[189761]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:118
	_go_fuzz_dep_.CoverTab[189762]++

										needsBase64 := true

										if protectedHeader != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:122
		_go_fuzz_dep_.CoverTab[189773]++
											var err error
											if needsBase64, err = protectedHeader.getB64(); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:124
			_go_fuzz_dep_.CoverTab[189774]++
												needsBase64 = true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:125
			// _ = "end of CoverTab[189774]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:126
			_go_fuzz_dep_.CoverTab[189775]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:126
			// _ = "end of CoverTab[189775]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:126
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:126
		// _ = "end of CoverTab[189773]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:127
		_go_fuzz_dep_.CoverTab[189776]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:127
		// _ = "end of CoverTab[189776]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:127
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:127
	// _ = "end of CoverTab[189762]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:127
	_go_fuzz_dep_.CoverTab[189763]++

										authData.WriteByte('.')

										if needsBase64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:131
		_go_fuzz_dep_.CoverTab[189777]++
											authData.WriteString(base64.RawURLEncoding.EncodeToString(payload))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:132
		// _ = "end of CoverTab[189777]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:133
		_go_fuzz_dep_.CoverTab[189778]++
											authData.Write(payload)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:134
		// _ = "end of CoverTab[189778]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:135
	// _ = "end of CoverTab[189763]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:135
	_go_fuzz_dep_.CoverTab[189764]++

										return authData.Bytes(), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:137
	// _ = "end of CoverTab[189764]"
}

// parseSignedFull parses a message in full format.
func parseSignedFull(input string) (*JSONWebSignature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:141
	_go_fuzz_dep_.CoverTab[189779]++
										var parsed rawJSONWebSignature
										err := json.Unmarshal([]byte(input), &parsed)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:144
		_go_fuzz_dep_.CoverTab[189781]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:145
		// _ = "end of CoverTab[189781]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:146
		_go_fuzz_dep_.CoverTab[189782]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:146
		// _ = "end of CoverTab[189782]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:146
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:146
	// _ = "end of CoverTab[189779]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:146
	_go_fuzz_dep_.CoverTab[189780]++

										return parsed.sanitized()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:148
	// _ = "end of CoverTab[189780]"
}

// sanitized produces a cleaned-up JWS object from the raw JSON.
func (parsed *rawJSONWebSignature) sanitized() (*JSONWebSignature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:152
	_go_fuzz_dep_.CoverTab[189783]++
										if parsed.Payload == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:153
		_go_fuzz_dep_.CoverTab[189787]++
											return nil, fmt.Errorf("square/go-jose: missing payload in JWS message")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:154
		// _ = "end of CoverTab[189787]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:155
		_go_fuzz_dep_.CoverTab[189788]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:155
		// _ = "end of CoverTab[189788]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:155
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:155
	// _ = "end of CoverTab[189783]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:155
	_go_fuzz_dep_.CoverTab[189784]++

										obj := &JSONWebSignature{
		payload:	parsed.Payload.bytes(),
		Signatures:	make([]Signature, len(parsed.Signatures)),
	}

	if len(parsed.Signatures) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:162
		_go_fuzz_dep_.CoverTab[189789]++

											signature := Signature{}
											if parsed.Protected != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:165
			_go_fuzz_dep_.CoverTab[189796]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:165
			return len(parsed.Protected.bytes()) > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:165
			// _ = "end of CoverTab[189796]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:165
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:165
			_go_fuzz_dep_.CoverTab[189797]++
												signature.protected = &rawHeader{}
												err := json.Unmarshal(parsed.Protected.bytes(), signature.protected)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:168
				_go_fuzz_dep_.CoverTab[189798]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:169
				// _ = "end of CoverTab[189798]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:170
				_go_fuzz_dep_.CoverTab[189799]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:170
				// _ = "end of CoverTab[189799]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:170
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:170
			// _ = "end of CoverTab[189797]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:171
			_go_fuzz_dep_.CoverTab[189800]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:171
			// _ = "end of CoverTab[189800]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:171
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:171
		// _ = "end of CoverTab[189789]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:171
		_go_fuzz_dep_.CoverTab[189790]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:174
		if parsed.Header != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:174
			_go_fuzz_dep_.CoverTab[189801]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:174
			return parsed.Header.getNonce() != ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:174
			// _ = "end of CoverTab[189801]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:174
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:174
			_go_fuzz_dep_.CoverTab[189802]++
												return nil, ErrUnprotectedNonce
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:175
			// _ = "end of CoverTab[189802]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:176
			_go_fuzz_dep_.CoverTab[189803]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:176
			// _ = "end of CoverTab[189803]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:176
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:176
		// _ = "end of CoverTab[189790]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:176
		_go_fuzz_dep_.CoverTab[189791]++

											signature.header = parsed.Header
											signature.Signature = parsed.Signature.bytes()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:189
		signature.original = &rawSignatureInfo{
			Protected:	parsed.Protected,
			Header:		parsed.Header,
			Signature:	parsed.Signature,
		}

		var err error
		signature.Header, err = signature.mergedHeaders().sanitized()
		if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:197
			_go_fuzz_dep_.CoverTab[189804]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:198
			// _ = "end of CoverTab[189804]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:199
			_go_fuzz_dep_.CoverTab[189805]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:199
			// _ = "end of CoverTab[189805]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:199
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:199
		// _ = "end of CoverTab[189791]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:199
		_go_fuzz_dep_.CoverTab[189792]++

											if signature.header != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:201
			_go_fuzz_dep_.CoverTab[189806]++
												signature.Unprotected, err = signature.header.sanitized()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:203
				_go_fuzz_dep_.CoverTab[189807]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:204
				// _ = "end of CoverTab[189807]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:205
				_go_fuzz_dep_.CoverTab[189808]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:205
				// _ = "end of CoverTab[189808]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:205
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:205
			// _ = "end of CoverTab[189806]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:206
			_go_fuzz_dep_.CoverTab[189809]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:206
			// _ = "end of CoverTab[189809]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:206
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:206
		// _ = "end of CoverTab[189792]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:206
		_go_fuzz_dep_.CoverTab[189793]++

											if signature.protected != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:208
			_go_fuzz_dep_.CoverTab[189810]++
												signature.Protected, err = signature.protected.sanitized()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:210
				_go_fuzz_dep_.CoverTab[189811]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:211
				// _ = "end of CoverTab[189811]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:212
				_go_fuzz_dep_.CoverTab[189812]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:212
				// _ = "end of CoverTab[189812]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:212
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:212
			// _ = "end of CoverTab[189810]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:213
			_go_fuzz_dep_.CoverTab[189813]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:213
			// _ = "end of CoverTab[189813]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:213
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:213
		// _ = "end of CoverTab[189793]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:213
		_go_fuzz_dep_.CoverTab[189794]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:216
		jwk := signature.Header.JSONWebKey
		if jwk != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:217
			_go_fuzz_dep_.CoverTab[189814]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:217
			return (!jwk.Valid() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:217
				_go_fuzz_dep_.CoverTab[189815]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:217
				return !jwk.IsPublic()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:217
				// _ = "end of CoverTab[189815]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:217
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:217
			// _ = "end of CoverTab[189814]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:217
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:217
			_go_fuzz_dep_.CoverTab[189816]++
												return nil, errors.New("square/go-jose: invalid embedded jwk, must be public key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:218
			// _ = "end of CoverTab[189816]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:219
			_go_fuzz_dep_.CoverTab[189817]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:219
			// _ = "end of CoverTab[189817]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:219
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:219
		// _ = "end of CoverTab[189794]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:219
		_go_fuzz_dep_.CoverTab[189795]++

											obj.Signatures = append(obj.Signatures, signature)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:221
		// _ = "end of CoverTab[189795]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:222
		_go_fuzz_dep_.CoverTab[189818]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:222
		// _ = "end of CoverTab[189818]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:222
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:222
	// _ = "end of CoverTab[189784]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:222
	_go_fuzz_dep_.CoverTab[189785]++

										for i, sig := range parsed.Signatures {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:224
		_go_fuzz_dep_.CoverTab[189819]++
											if sig.Protected != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:225
			_go_fuzz_dep_.CoverTab[189826]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:225
			return len(sig.Protected.bytes()) > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:225
			// _ = "end of CoverTab[189826]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:225
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:225
			_go_fuzz_dep_.CoverTab[189827]++
												obj.Signatures[i].protected = &rawHeader{}
												err := json.Unmarshal(sig.Protected.bytes(), obj.Signatures[i].protected)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:228
				_go_fuzz_dep_.CoverTab[189828]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:229
				// _ = "end of CoverTab[189828]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:230
				_go_fuzz_dep_.CoverTab[189829]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:230
				// _ = "end of CoverTab[189829]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:230
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:230
			// _ = "end of CoverTab[189827]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:231
			_go_fuzz_dep_.CoverTab[189830]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:231
			// _ = "end of CoverTab[189830]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:231
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:231
		// _ = "end of CoverTab[189819]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:231
		_go_fuzz_dep_.CoverTab[189820]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:234
		if sig.Header != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:234
			_go_fuzz_dep_.CoverTab[189831]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:234
			return sig.Header.getNonce() != ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:234
			// _ = "end of CoverTab[189831]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:234
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:234
			_go_fuzz_dep_.CoverTab[189832]++
												return nil, ErrUnprotectedNonce
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:235
			// _ = "end of CoverTab[189832]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:236
			_go_fuzz_dep_.CoverTab[189833]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:236
			// _ = "end of CoverTab[189833]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:236
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:236
		// _ = "end of CoverTab[189820]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:236
		_go_fuzz_dep_.CoverTab[189821]++

											var err error
											obj.Signatures[i].Header, err = obj.Signatures[i].mergedHeaders().sanitized()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:240
			_go_fuzz_dep_.CoverTab[189834]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:241
			// _ = "end of CoverTab[189834]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:242
			_go_fuzz_dep_.CoverTab[189835]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:242
			// _ = "end of CoverTab[189835]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:242
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:242
		// _ = "end of CoverTab[189821]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:242
		_go_fuzz_dep_.CoverTab[189822]++

											if obj.Signatures[i].header != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:244
			_go_fuzz_dep_.CoverTab[189836]++
												obj.Signatures[i].Unprotected, err = obj.Signatures[i].header.sanitized()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:246
				_go_fuzz_dep_.CoverTab[189837]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:247
				// _ = "end of CoverTab[189837]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:248
				_go_fuzz_dep_.CoverTab[189838]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:248
				// _ = "end of CoverTab[189838]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:248
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:248
			// _ = "end of CoverTab[189836]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:249
			_go_fuzz_dep_.CoverTab[189839]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:249
			// _ = "end of CoverTab[189839]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:249
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:249
		// _ = "end of CoverTab[189822]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:249
		_go_fuzz_dep_.CoverTab[189823]++

											if obj.Signatures[i].protected != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:251
			_go_fuzz_dep_.CoverTab[189840]++
												obj.Signatures[i].Protected, err = obj.Signatures[i].protected.sanitized()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:253
				_go_fuzz_dep_.CoverTab[189841]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:254
				// _ = "end of CoverTab[189841]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:255
				_go_fuzz_dep_.CoverTab[189842]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:255
				// _ = "end of CoverTab[189842]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:255
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:255
			// _ = "end of CoverTab[189840]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:256
			_go_fuzz_dep_.CoverTab[189843]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:256
			// _ = "end of CoverTab[189843]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:256
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:256
		// _ = "end of CoverTab[189823]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:256
		_go_fuzz_dep_.CoverTab[189824]++

											obj.Signatures[i].Signature = sig.Signature.bytes()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:261
		jwk := obj.Signatures[i].Header.JSONWebKey
		if jwk != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:262
			_go_fuzz_dep_.CoverTab[189844]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:262
			return (!jwk.Valid() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:262
				_go_fuzz_dep_.CoverTab[189845]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:262
				return !jwk.IsPublic()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:262
				// _ = "end of CoverTab[189845]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:262
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:262
			// _ = "end of CoverTab[189844]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:262
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:262
			_go_fuzz_dep_.CoverTab[189846]++
												return nil, errors.New("square/go-jose: invalid embedded jwk, must be public key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:263
			// _ = "end of CoverTab[189846]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:264
			_go_fuzz_dep_.CoverTab[189847]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:264
			// _ = "end of CoverTab[189847]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:264
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:264
		// _ = "end of CoverTab[189824]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:264
		_go_fuzz_dep_.CoverTab[189825]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:267
		original := sig

											obj.Signatures[i].header = sig.Header
											obj.Signatures[i].original = &original
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:270
		// _ = "end of CoverTab[189825]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:271
	// _ = "end of CoverTab[189785]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:271
	_go_fuzz_dep_.CoverTab[189786]++

										return obj, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:273
	// _ = "end of CoverTab[189786]"
}

// parseSignedCompact parses a message in compact format.
func parseSignedCompact(input string, payload []byte) (*JSONWebSignature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:277
	_go_fuzz_dep_.CoverTab[189848]++
										parts := strings.Split(input, ".")
										if len(parts) != 3 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:279
		_go_fuzz_dep_.CoverTab[189854]++
											return nil, fmt.Errorf("square/go-jose: compact JWS format must have three parts")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:280
		// _ = "end of CoverTab[189854]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:281
		_go_fuzz_dep_.CoverTab[189855]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:281
		// _ = "end of CoverTab[189855]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:281
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:281
	// _ = "end of CoverTab[189848]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:281
	_go_fuzz_dep_.CoverTab[189849]++

										if parts[1] != "" && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:283
		_go_fuzz_dep_.CoverTab[189856]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:283
		return payload != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:283
		// _ = "end of CoverTab[189856]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:283
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:283
		_go_fuzz_dep_.CoverTab[189857]++
											return nil, fmt.Errorf("square/go-jose: payload is not detached")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:284
		// _ = "end of CoverTab[189857]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:285
		_go_fuzz_dep_.CoverTab[189858]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:285
		// _ = "end of CoverTab[189858]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:285
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:285
	// _ = "end of CoverTab[189849]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:285
	_go_fuzz_dep_.CoverTab[189850]++

										rawProtected, err := base64.RawURLEncoding.DecodeString(parts[0])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:288
		_go_fuzz_dep_.CoverTab[189859]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:289
		// _ = "end of CoverTab[189859]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:290
		_go_fuzz_dep_.CoverTab[189860]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:290
		// _ = "end of CoverTab[189860]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:290
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:290
	// _ = "end of CoverTab[189850]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:290
	_go_fuzz_dep_.CoverTab[189851]++

										if payload == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:292
		_go_fuzz_dep_.CoverTab[189861]++
											payload, err = base64.RawURLEncoding.DecodeString(parts[1])
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:294
			_go_fuzz_dep_.CoverTab[189862]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:295
			// _ = "end of CoverTab[189862]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:296
			_go_fuzz_dep_.CoverTab[189863]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:296
			// _ = "end of CoverTab[189863]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:296
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:296
		// _ = "end of CoverTab[189861]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:297
		_go_fuzz_dep_.CoverTab[189864]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:297
		// _ = "end of CoverTab[189864]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:297
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:297
	// _ = "end of CoverTab[189851]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:297
	_go_fuzz_dep_.CoverTab[189852]++

										signature, err := base64.RawURLEncoding.DecodeString(parts[2])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:300
		_go_fuzz_dep_.CoverTab[189865]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:301
		// _ = "end of CoverTab[189865]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:302
		_go_fuzz_dep_.CoverTab[189866]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:302
		// _ = "end of CoverTab[189866]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:302
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:302
	// _ = "end of CoverTab[189852]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:302
	_go_fuzz_dep_.CoverTab[189853]++

										raw := &rawJSONWebSignature{
		Payload:	newBuffer(payload),
		Protected:	newBuffer(rawProtected),
		Signature:	newBuffer(signature),
	}
										return raw.sanitized()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:309
	// _ = "end of CoverTab[189853]"
}

func (obj JSONWebSignature) compactSerialize(detached bool) (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:312
	_go_fuzz_dep_.CoverTab[189867]++
										if len(obj.Signatures) != 1 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:313
		_go_fuzz_dep_.CoverTab[189870]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:313
		return obj.Signatures[0].header != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:313
		// _ = "end of CoverTab[189870]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:313
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:313
		_go_fuzz_dep_.CoverTab[189871]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:313
		return obj.Signatures[0].protected == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:313
		// _ = "end of CoverTab[189871]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:313
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:313
		_go_fuzz_dep_.CoverTab[189872]++
											return "", ErrNotSupported
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:314
		// _ = "end of CoverTab[189872]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:315
		_go_fuzz_dep_.CoverTab[189873]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:315
		// _ = "end of CoverTab[189873]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:315
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:315
	// _ = "end of CoverTab[189867]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:315
	_go_fuzz_dep_.CoverTab[189868]++

										serializedProtected := base64.RawURLEncoding.EncodeToString(mustSerializeJSON(obj.Signatures[0].protected))
										payload := ""
										signature := base64.RawURLEncoding.EncodeToString(obj.Signatures[0].Signature)

										if !detached {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:321
		_go_fuzz_dep_.CoverTab[189874]++
											payload = base64.RawURLEncoding.EncodeToString(obj.payload)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:322
		// _ = "end of CoverTab[189874]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:323
		_go_fuzz_dep_.CoverTab[189875]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:323
		// _ = "end of CoverTab[189875]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:323
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:323
	// _ = "end of CoverTab[189868]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:323
	_go_fuzz_dep_.CoverTab[189869]++

										return fmt.Sprintf("%s.%s.%s", serializedProtected, payload, signature), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:325
	// _ = "end of CoverTab[189869]"
}

// CompactSerialize serializes an object using the compact serialization format.
func (obj JSONWebSignature) CompactSerialize() (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:329
	_go_fuzz_dep_.CoverTab[189876]++
										return obj.compactSerialize(false)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:330
	// _ = "end of CoverTab[189876]"
}

// DetachedCompactSerialize serializes an object using the compact serialization format with detached payload.
func (obj JSONWebSignature) DetachedCompactSerialize() (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:334
	_go_fuzz_dep_.CoverTab[189877]++
										return obj.compactSerialize(true)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:335
	// _ = "end of CoverTab[189877]"
}

// FullSerialize serializes an object using the full JSON serialization format.
func (obj JSONWebSignature) FullSerialize() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:339
	_go_fuzz_dep_.CoverTab[189878]++
										raw := rawJSONWebSignature{
		Payload: newBuffer(obj.payload),
	}

	if len(obj.Signatures) == 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:344
		_go_fuzz_dep_.CoverTab[189880]++
											if obj.Signatures[0].protected != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:345
			_go_fuzz_dep_.CoverTab[189882]++
												serializedProtected := mustSerializeJSON(obj.Signatures[0].protected)
												raw.Protected = newBuffer(serializedProtected)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:347
			// _ = "end of CoverTab[189882]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:348
			_go_fuzz_dep_.CoverTab[189883]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:348
			// _ = "end of CoverTab[189883]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:348
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:348
		// _ = "end of CoverTab[189880]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:348
		_go_fuzz_dep_.CoverTab[189881]++
											raw.Header = obj.Signatures[0].header
											raw.Signature = newBuffer(obj.Signatures[0].Signature)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:350
		// _ = "end of CoverTab[189881]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:351
		_go_fuzz_dep_.CoverTab[189884]++
											raw.Signatures = make([]rawSignatureInfo, len(obj.Signatures))
											for i, signature := range obj.Signatures {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:353
			_go_fuzz_dep_.CoverTab[189885]++
												raw.Signatures[i] = rawSignatureInfo{
				Header:		signature.header,
				Signature:	newBuffer(signature.Signature),
			}

			if signature.protected != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:359
				_go_fuzz_dep_.CoverTab[189886]++
													raw.Signatures[i].Protected = newBuffer(mustSerializeJSON(signature.protected))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:360
				// _ = "end of CoverTab[189886]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:361
				_go_fuzz_dep_.CoverTab[189887]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:361
				// _ = "end of CoverTab[189887]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:361
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:361
			// _ = "end of CoverTab[189885]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:362
		// _ = "end of CoverTab[189884]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:363
	// _ = "end of CoverTab[189878]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:363
	_go_fuzz_dep_.CoverTab[189879]++

										return string(mustSerializeJSON(raw))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:365
	// _ = "end of CoverTab[189879]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:366
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jws.go:366
var _ = _go_fuzz_dep_.CoverTab
