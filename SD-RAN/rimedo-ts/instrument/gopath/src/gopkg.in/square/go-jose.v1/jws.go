//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:17
)

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/square/go-jose.v1/json"
)

// rawJsonWebSignature represents a raw JWS JSON object. Used for parsing/serializing.
type rawJsonWebSignature struct {
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

// JsonWebSignature represents a signed JWS object after parsing.
type JsonWebSignature struct {
	payload	[]byte
	// Signatures attached to this object (may be more than one for multi-sig).
	// Be careful about accessing these directly, prefer to use Verify() or
	// VerifyMulti() to ensure that the data you're getting is verified.
	Signatures	[]Signature
}

// Signature represents a single signature over the JWS payload and protected header.
type Signature struct {
	// Header fields, such as the signature algorithm
	Header	JoseHeader

	// The actual signature value
	Signature	[]byte

	protected	*rawHeader
	header		*rawHeader
	original	*rawSignatureInfo
}

// ParseSigned parses a signed message in compact or full serialization format.
func ParseSigned(input string) (*JsonWebSignature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:66
	_go_fuzz_dep_.CoverTab[186447]++
										input = stripWhitespace(input)
										if strings.HasPrefix(input, "{") {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:68
		_go_fuzz_dep_.CoverTab[186449]++
											return parseSignedFull(input)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:69
		// _ = "end of CoverTab[186449]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:70
		_go_fuzz_dep_.CoverTab[186450]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:70
		// _ = "end of CoverTab[186450]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:70
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:70
	// _ = "end of CoverTab[186447]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:70
	_go_fuzz_dep_.CoverTab[186448]++

										return parseSignedCompact(input)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:72
	// _ = "end of CoverTab[186448]"
}

// Get a header value
func (sig Signature) mergedHeaders() rawHeader {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:76
	_go_fuzz_dep_.CoverTab[186451]++
										out := rawHeader{}
										out.merge(sig.protected)
										out.merge(sig.header)
										return out
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:80
	// _ = "end of CoverTab[186451]"
}

// Compute data to be signed
func (obj JsonWebSignature) computeAuthData(signature *Signature) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:84
	_go_fuzz_dep_.CoverTab[186452]++
										var serializedProtected string

										if signature.original != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:87
		_go_fuzz_dep_.CoverTab[186454]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:87
		return signature.original.Protected != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:87
		// _ = "end of CoverTab[186454]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:87
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:87
		_go_fuzz_dep_.CoverTab[186455]++
											serializedProtected = signature.original.Protected.base64()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:88
		// _ = "end of CoverTab[186455]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:89
		_go_fuzz_dep_.CoverTab[186456]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:89
		if signature.protected != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:89
			_go_fuzz_dep_.CoverTab[186457]++
												serializedProtected = base64URLEncode(mustSerializeJSON(signature.protected))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:90
			// _ = "end of CoverTab[186457]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:91
			_go_fuzz_dep_.CoverTab[186458]++
												serializedProtected = ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:92
			// _ = "end of CoverTab[186458]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:93
		// _ = "end of CoverTab[186456]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:93
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:93
	// _ = "end of CoverTab[186452]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:93
	_go_fuzz_dep_.CoverTab[186453]++

										return []byte(fmt.Sprintf("%s.%s",
		serializedProtected,
		base64URLEncode(obj.payload)))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:97
	// _ = "end of CoverTab[186453]"
}

// parseSignedFull parses a message in full format.
func parseSignedFull(input string) (*JsonWebSignature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:101
	_go_fuzz_dep_.CoverTab[186459]++
										var parsed rawJsonWebSignature
										err := json.Unmarshal([]byte(input), &parsed)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:104
		_go_fuzz_dep_.CoverTab[186461]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:105
		// _ = "end of CoverTab[186461]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:106
		_go_fuzz_dep_.CoverTab[186462]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:106
		// _ = "end of CoverTab[186462]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:106
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:106
	// _ = "end of CoverTab[186459]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:106
	_go_fuzz_dep_.CoverTab[186460]++

										return parsed.sanitized()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:108
	// _ = "end of CoverTab[186460]"
}

// sanitized produces a cleaned-up JWS object from the raw JSON.
func (parsed *rawJsonWebSignature) sanitized() (*JsonWebSignature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:112
	_go_fuzz_dep_.CoverTab[186463]++
										if parsed.Payload == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:113
		_go_fuzz_dep_.CoverTab[186467]++
											return nil, fmt.Errorf("square/go-jose: missing payload in JWS message")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:114
		// _ = "end of CoverTab[186467]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:115
		_go_fuzz_dep_.CoverTab[186468]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:115
		// _ = "end of CoverTab[186468]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:115
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:115
	// _ = "end of CoverTab[186463]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:115
	_go_fuzz_dep_.CoverTab[186464]++

										obj := &JsonWebSignature{
		payload:	parsed.Payload.bytes(),
		Signatures:	make([]Signature, len(parsed.Signatures)),
	}

	if len(parsed.Signatures) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:122
		_go_fuzz_dep_.CoverTab[186469]++

											signature := Signature{}
											if parsed.Protected != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:125
			_go_fuzz_dep_.CoverTab[186473]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:125
			return len(parsed.Protected.bytes()) > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:125
			// _ = "end of CoverTab[186473]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:125
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:125
			_go_fuzz_dep_.CoverTab[186474]++
												signature.protected = &rawHeader{}
												err := json.Unmarshal(parsed.Protected.bytes(), signature.protected)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:128
				_go_fuzz_dep_.CoverTab[186475]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:129
				// _ = "end of CoverTab[186475]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:130
				_go_fuzz_dep_.CoverTab[186476]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:130
				// _ = "end of CoverTab[186476]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:130
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:130
			// _ = "end of CoverTab[186474]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:131
			_go_fuzz_dep_.CoverTab[186477]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:131
			// _ = "end of CoverTab[186477]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:131
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:131
		// _ = "end of CoverTab[186469]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:131
		_go_fuzz_dep_.CoverTab[186470]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:134
		if parsed.Header != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:134
			_go_fuzz_dep_.CoverTab[186478]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:134
			return parsed.Header.Nonce != ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:134
			// _ = "end of CoverTab[186478]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:134
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:134
			_go_fuzz_dep_.CoverTab[186479]++
												return nil, ErrUnprotectedNonce
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:135
			// _ = "end of CoverTab[186479]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:136
			_go_fuzz_dep_.CoverTab[186480]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:136
			// _ = "end of CoverTab[186480]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:136
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:136
		// _ = "end of CoverTab[186470]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:136
		_go_fuzz_dep_.CoverTab[186471]++

											signature.header = parsed.Header
											signature.Signature = parsed.Signature.bytes()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:149
		signature.original = &rawSignatureInfo{
			Protected:	parsed.Protected,
			Header:		parsed.Header,
			Signature:	parsed.Signature,
		}

											signature.Header = signature.mergedHeaders().sanitized()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:158
		jwk := signature.Header.JsonWebKey
		if jwk != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:159
			_go_fuzz_dep_.CoverTab[186481]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:159
			return (!jwk.Valid() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:159
				_go_fuzz_dep_.CoverTab[186482]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:159
				return !jwk.IsPublic()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:159
				// _ = "end of CoverTab[186482]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:159
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:159
			// _ = "end of CoverTab[186481]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:159
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:159
			_go_fuzz_dep_.CoverTab[186483]++
												return nil, errors.New("square/go-jose: invalid embedded jwk, must be public key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:160
			// _ = "end of CoverTab[186483]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:161
			_go_fuzz_dep_.CoverTab[186484]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:161
			// _ = "end of CoverTab[186484]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:161
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:161
		// _ = "end of CoverTab[186471]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:161
		_go_fuzz_dep_.CoverTab[186472]++

											obj.Signatures = append(obj.Signatures, signature)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:163
		// _ = "end of CoverTab[186472]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:164
		_go_fuzz_dep_.CoverTab[186485]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:164
		// _ = "end of CoverTab[186485]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:164
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:164
	// _ = "end of CoverTab[186464]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:164
	_go_fuzz_dep_.CoverTab[186465]++

										for i, sig := range parsed.Signatures {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:166
		_go_fuzz_dep_.CoverTab[186486]++
											if sig.Protected != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:167
			_go_fuzz_dep_.CoverTab[186490]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:167
			return len(sig.Protected.bytes()) > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:167
			// _ = "end of CoverTab[186490]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:167
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:167
			_go_fuzz_dep_.CoverTab[186491]++
												obj.Signatures[i].protected = &rawHeader{}
												err := json.Unmarshal(sig.Protected.bytes(), obj.Signatures[i].protected)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:170
				_go_fuzz_dep_.CoverTab[186492]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:171
				// _ = "end of CoverTab[186492]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:172
				_go_fuzz_dep_.CoverTab[186493]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:172
				// _ = "end of CoverTab[186493]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:172
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:172
			// _ = "end of CoverTab[186491]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:173
			_go_fuzz_dep_.CoverTab[186494]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:173
			// _ = "end of CoverTab[186494]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:173
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:173
		// _ = "end of CoverTab[186486]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:173
		_go_fuzz_dep_.CoverTab[186487]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:176
		if sig.Header != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:176
			_go_fuzz_dep_.CoverTab[186495]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:176
			return sig.Header.Nonce != ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:176
			// _ = "end of CoverTab[186495]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:176
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:176
			_go_fuzz_dep_.CoverTab[186496]++
												return nil, ErrUnprotectedNonce
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:177
			// _ = "end of CoverTab[186496]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:178
			_go_fuzz_dep_.CoverTab[186497]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:178
			// _ = "end of CoverTab[186497]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:178
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:178
		// _ = "end of CoverTab[186487]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:178
		_go_fuzz_dep_.CoverTab[186488]++

											obj.Signatures[i].Header = obj.Signatures[i].mergedHeaders().sanitized()
											obj.Signatures[i].Signature = sig.Signature.bytes()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:184
		jwk := obj.Signatures[i].Header.JsonWebKey
		if jwk != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:185
			_go_fuzz_dep_.CoverTab[186498]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:185
			return (!jwk.Valid() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:185
				_go_fuzz_dep_.CoverTab[186499]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:185
				return !jwk.IsPublic()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:185
				// _ = "end of CoverTab[186499]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:185
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:185
			// _ = "end of CoverTab[186498]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:185
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:185
			_go_fuzz_dep_.CoverTab[186500]++
												return nil, errors.New("square/go-jose: invalid embedded jwk, must be public key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:186
			// _ = "end of CoverTab[186500]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:187
			_go_fuzz_dep_.CoverTab[186501]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:187
			// _ = "end of CoverTab[186501]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:187
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:187
		// _ = "end of CoverTab[186488]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:187
		_go_fuzz_dep_.CoverTab[186489]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:190
		original := sig

											obj.Signatures[i].header = sig.Header
											obj.Signatures[i].original = &original
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:193
		// _ = "end of CoverTab[186489]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:194
	// _ = "end of CoverTab[186465]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:194
	_go_fuzz_dep_.CoverTab[186466]++

										return obj, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:196
	// _ = "end of CoverTab[186466]"
}

// parseSignedCompact parses a message in compact format.
func parseSignedCompact(input string) (*JsonWebSignature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:200
	_go_fuzz_dep_.CoverTab[186502]++
										parts := strings.Split(input, ".")
										if len(parts) != 3 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:202
		_go_fuzz_dep_.CoverTab[186507]++
											return nil, fmt.Errorf("square/go-jose: compact JWS format must have three parts")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:203
		// _ = "end of CoverTab[186507]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:204
		_go_fuzz_dep_.CoverTab[186508]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:204
		// _ = "end of CoverTab[186508]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:204
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:204
	// _ = "end of CoverTab[186502]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:204
	_go_fuzz_dep_.CoverTab[186503]++

										rawProtected, err := base64URLDecode(parts[0])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:207
		_go_fuzz_dep_.CoverTab[186509]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:208
		// _ = "end of CoverTab[186509]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:209
		_go_fuzz_dep_.CoverTab[186510]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:209
		// _ = "end of CoverTab[186510]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:209
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:209
	// _ = "end of CoverTab[186503]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:209
	_go_fuzz_dep_.CoverTab[186504]++

										payload, err := base64URLDecode(parts[1])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:212
		_go_fuzz_dep_.CoverTab[186511]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:213
		// _ = "end of CoverTab[186511]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:214
		_go_fuzz_dep_.CoverTab[186512]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:214
		// _ = "end of CoverTab[186512]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:214
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:214
	// _ = "end of CoverTab[186504]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:214
	_go_fuzz_dep_.CoverTab[186505]++

										signature, err := base64URLDecode(parts[2])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:217
		_go_fuzz_dep_.CoverTab[186513]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:218
		// _ = "end of CoverTab[186513]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:219
		_go_fuzz_dep_.CoverTab[186514]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:219
		// _ = "end of CoverTab[186514]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:219
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:219
	// _ = "end of CoverTab[186505]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:219
	_go_fuzz_dep_.CoverTab[186506]++

										raw := &rawJsonWebSignature{
		Payload:	newBuffer(payload),
		Protected:	newBuffer(rawProtected),
		Signature:	newBuffer(signature),
	}
										return raw.sanitized()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:226
	// _ = "end of CoverTab[186506]"
}

// CompactSerialize serializes an object using the compact serialization format.
func (obj JsonWebSignature) CompactSerialize() (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:230
	_go_fuzz_dep_.CoverTab[186515]++
										if len(obj.Signatures) != 1 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:231
		_go_fuzz_dep_.CoverTab[186517]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:231
		return obj.Signatures[0].header != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:231
		// _ = "end of CoverTab[186517]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:231
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:231
		_go_fuzz_dep_.CoverTab[186518]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:231
		return obj.Signatures[0].protected == nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:231
		// _ = "end of CoverTab[186518]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:231
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:231
		_go_fuzz_dep_.CoverTab[186519]++
											return "", ErrNotSupported
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:232
		// _ = "end of CoverTab[186519]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:233
		_go_fuzz_dep_.CoverTab[186520]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:233
		// _ = "end of CoverTab[186520]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:233
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:233
	// _ = "end of CoverTab[186515]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:233
	_go_fuzz_dep_.CoverTab[186516]++

										serializedProtected := mustSerializeJSON(obj.Signatures[0].protected)

										return fmt.Sprintf(
		"%s.%s.%s",
		base64URLEncode(serializedProtected),
		base64URLEncode(obj.payload),
		base64URLEncode(obj.Signatures[0].Signature)), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:241
	// _ = "end of CoverTab[186516]"
}

// FullSerialize serializes an object using the full JSON serialization format.
func (obj JsonWebSignature) FullSerialize() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:245
	_go_fuzz_dep_.CoverTab[186521]++
										raw := rawJsonWebSignature{
		Payload: newBuffer(obj.payload),
	}

	if len(obj.Signatures) == 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:250
		_go_fuzz_dep_.CoverTab[186523]++
											if obj.Signatures[0].protected != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:251
			_go_fuzz_dep_.CoverTab[186525]++
												serializedProtected := mustSerializeJSON(obj.Signatures[0].protected)
												raw.Protected = newBuffer(serializedProtected)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:253
			// _ = "end of CoverTab[186525]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:254
			_go_fuzz_dep_.CoverTab[186526]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:254
			// _ = "end of CoverTab[186526]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:254
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:254
		// _ = "end of CoverTab[186523]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:254
		_go_fuzz_dep_.CoverTab[186524]++
											raw.Header = obj.Signatures[0].header
											raw.Signature = newBuffer(obj.Signatures[0].Signature)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:256
		// _ = "end of CoverTab[186524]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:257
		_go_fuzz_dep_.CoverTab[186527]++
											raw.Signatures = make([]rawSignatureInfo, len(obj.Signatures))
											for i, signature := range obj.Signatures {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:259
			_go_fuzz_dep_.CoverTab[186528]++
												raw.Signatures[i] = rawSignatureInfo{
				Header:		signature.header,
				Signature:	newBuffer(signature.Signature),
			}

			if signature.protected != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:265
				_go_fuzz_dep_.CoverTab[186529]++
													raw.Signatures[i].Protected = newBuffer(mustSerializeJSON(signature.protected))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:266
				// _ = "end of CoverTab[186529]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:267
				_go_fuzz_dep_.CoverTab[186530]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:267
				// _ = "end of CoverTab[186530]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:267
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:267
			// _ = "end of CoverTab[186528]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:268
		// _ = "end of CoverTab[186527]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:269
	// _ = "end of CoverTab[186521]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:269
	_go_fuzz_dep_.CoverTab[186522]++

										return string(mustSerializeJSON(raw))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:271
	// _ = "end of CoverTab[186522]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:272
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jws.go:272
var _ = _go_fuzz_dep_.CoverTab
