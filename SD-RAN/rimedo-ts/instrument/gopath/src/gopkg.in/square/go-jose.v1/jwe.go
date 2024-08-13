//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:17
)

import (
	"fmt"
	"strings"

	"gopkg.in/square/go-jose.v1/json"
)

// rawJsonWebEncryption represents a raw JWE JSON object. Used for parsing/serializing.
type rawJsonWebEncryption struct {
	Protected	*byteBuffer		`json:"protected,omitempty"`
	Unprotected	*rawHeader		`json:"unprotected,omitempty"`
	Header		*rawHeader		`json:"header,omitempty"`
	Recipients	[]rawRecipientInfo	`json:"recipients,omitempty"`
	Aad		*byteBuffer		`json:"aad,omitempty"`
	EncryptedKey	*byteBuffer		`json:"encrypted_key,omitempty"`
	Iv		*byteBuffer		`json:"iv,omitempty"`
	Ciphertext	*byteBuffer		`json:"ciphertext,omitempty"`
	Tag		*byteBuffer		`json:"tag,omitempty"`
}

// rawRecipientInfo represents a raw JWE Per-Recipient header JSON object. Used for parsing/serializing.
type rawRecipientInfo struct {
	Header		*rawHeader	`json:"header,omitempty"`
	EncryptedKey	string		`json:"encrypted_key,omitempty"`
}

// JsonWebEncryption represents an encrypted JWE object after parsing.
type JsonWebEncryption struct {
	Header				JoseHeader
	protected, unprotected		*rawHeader
	recipients			[]recipientInfo
	aad, iv, ciphertext, tag	[]byte
	original			*rawJsonWebEncryption
}

// recipientInfo represents a raw JWE Per-Recipient header JSON object after parsing.
type recipientInfo struct {
	header		*rawHeader
	encryptedKey	[]byte
}

// GetAuthData retrieves the (optional) authenticated data attached to the object.
func (obj JsonWebEncryption) GetAuthData() []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:61
	_go_fuzz_dep_.CoverTab[186186]++
										if obj.aad != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:62
		_go_fuzz_dep_.CoverTab[186188]++
											out := make([]byte, len(obj.aad))
											copy(out, obj.aad)
											return out
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:65
		// _ = "end of CoverTab[186188]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:66
		_go_fuzz_dep_.CoverTab[186189]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:66
		// _ = "end of CoverTab[186189]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:66
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:66
	// _ = "end of CoverTab[186186]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:66
	_go_fuzz_dep_.CoverTab[186187]++

										return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:68
	// _ = "end of CoverTab[186187]"
}

// Get the merged header values
func (obj JsonWebEncryption) mergedHeaders(recipient *recipientInfo) rawHeader {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:72
	_go_fuzz_dep_.CoverTab[186190]++
										out := rawHeader{}
										out.merge(obj.protected)
										out.merge(obj.unprotected)

										if recipient != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:77
		_go_fuzz_dep_.CoverTab[186192]++
											out.merge(recipient.header)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:78
		// _ = "end of CoverTab[186192]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:79
		_go_fuzz_dep_.CoverTab[186193]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:79
		// _ = "end of CoverTab[186193]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:79
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:79
	// _ = "end of CoverTab[186190]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:79
	_go_fuzz_dep_.CoverTab[186191]++

										return out
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:81
	// _ = "end of CoverTab[186191]"
}

// Get the additional authenticated data from a JWE object.
func (obj JsonWebEncryption) computeAuthData() []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:85
	_go_fuzz_dep_.CoverTab[186194]++
										var protected string

										if obj.original != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:88
		_go_fuzz_dep_.CoverTab[186197]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:88
		return obj.original.Protected != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:88
		// _ = "end of CoverTab[186197]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:88
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:88
		_go_fuzz_dep_.CoverTab[186198]++
											protected = obj.original.Protected.base64()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:89
		// _ = "end of CoverTab[186198]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:90
		_go_fuzz_dep_.CoverTab[186199]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:90
		if obj.protected != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:90
			_go_fuzz_dep_.CoverTab[186200]++
												protected = base64URLEncode(mustSerializeJSON((obj.protected)))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:91
			// _ = "end of CoverTab[186200]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:92
			_go_fuzz_dep_.CoverTab[186201]++
												protected = ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:93
			// _ = "end of CoverTab[186201]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:94
		// _ = "end of CoverTab[186199]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:94
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:94
	// _ = "end of CoverTab[186194]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:94
	_go_fuzz_dep_.CoverTab[186195]++

										output := []byte(protected)
										if obj.aad != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:97
		_go_fuzz_dep_.CoverTab[186202]++
											output = append(output, '.')
											output = append(output, []byte(base64URLEncode(obj.aad))...)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:99
		// _ = "end of CoverTab[186202]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:100
		_go_fuzz_dep_.CoverTab[186203]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:100
		// _ = "end of CoverTab[186203]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:100
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:100
	// _ = "end of CoverTab[186195]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:100
	_go_fuzz_dep_.CoverTab[186196]++

										return output
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:102
	// _ = "end of CoverTab[186196]"
}

// ParseEncrypted parses an encrypted message in compact or full serialization format.
func ParseEncrypted(input string) (*JsonWebEncryption, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:106
	_go_fuzz_dep_.CoverTab[186204]++
										input = stripWhitespace(input)
										if strings.HasPrefix(input, "{") {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:108
		_go_fuzz_dep_.CoverTab[186206]++
											return parseEncryptedFull(input)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:109
		// _ = "end of CoverTab[186206]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:110
		_go_fuzz_dep_.CoverTab[186207]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:110
		// _ = "end of CoverTab[186207]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:110
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:110
	// _ = "end of CoverTab[186204]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:110
	_go_fuzz_dep_.CoverTab[186205]++

										return parseEncryptedCompact(input)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:112
	// _ = "end of CoverTab[186205]"
}

// parseEncryptedFull parses a message in compact format.
func parseEncryptedFull(input string) (*JsonWebEncryption, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:116
	_go_fuzz_dep_.CoverTab[186208]++
										var parsed rawJsonWebEncryption
										err := json.Unmarshal([]byte(input), &parsed)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:119
		_go_fuzz_dep_.CoverTab[186210]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:120
		// _ = "end of CoverTab[186210]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:121
		_go_fuzz_dep_.CoverTab[186211]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:121
		// _ = "end of CoverTab[186211]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:121
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:121
	// _ = "end of CoverTab[186208]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:121
	_go_fuzz_dep_.CoverTab[186209]++

										return parsed.sanitized()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:123
	// _ = "end of CoverTab[186209]"
}

// sanitized produces a cleaned-up JWE object from the raw JSON.
func (parsed *rawJsonWebEncryption) sanitized() (*JsonWebEncryption, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:127
	_go_fuzz_dep_.CoverTab[186212]++
										obj := &JsonWebEncryption{
		original:	parsed,
		unprotected:	parsed.Unprotected,
	}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:134
	if (parsed.Unprotected != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:134
		_go_fuzz_dep_.CoverTab[186217]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:134
		return parsed.Unprotected.Nonce != ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:134
		// _ = "end of CoverTab[186217]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:134
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:134
		_go_fuzz_dep_.CoverTab[186218]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:134
		return (parsed.Header != nil && func() bool {
												_go_fuzz_dep_.CoverTab[186219]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:135
			return parsed.Header.Nonce != ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:135
			// _ = "end of CoverTab[186219]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:135
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:135
		// _ = "end of CoverTab[186218]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:135
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:135
		_go_fuzz_dep_.CoverTab[186220]++
											return nil, ErrUnprotectedNonce
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:136
		// _ = "end of CoverTab[186220]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:137
		_go_fuzz_dep_.CoverTab[186221]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:137
		// _ = "end of CoverTab[186221]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:137
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:137
	// _ = "end of CoverTab[186212]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:137
	_go_fuzz_dep_.CoverTab[186213]++

										if parsed.Protected != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:139
		_go_fuzz_dep_.CoverTab[186222]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:139
		return len(parsed.Protected.bytes()) > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:139
		// _ = "end of CoverTab[186222]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:139
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:139
		_go_fuzz_dep_.CoverTab[186223]++
											err := json.Unmarshal(parsed.Protected.bytes(), &obj.protected)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:141
			_go_fuzz_dep_.CoverTab[186224]++
												return nil, fmt.Errorf("square/go-jose: invalid protected header: %s, %s", err, parsed.Protected.base64())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:142
			// _ = "end of CoverTab[186224]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:143
			_go_fuzz_dep_.CoverTab[186225]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:143
			// _ = "end of CoverTab[186225]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:143
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:143
		// _ = "end of CoverTab[186223]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:144
		_go_fuzz_dep_.CoverTab[186226]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:144
		// _ = "end of CoverTab[186226]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:144
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:144
	// _ = "end of CoverTab[186213]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:144
	_go_fuzz_dep_.CoverTab[186214]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:148
	obj.Header = obj.mergedHeaders(nil).sanitized()

	if len(parsed.Recipients) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:150
		_go_fuzz_dep_.CoverTab[186227]++
											obj.recipients = []recipientInfo{
			recipientInfo{
				header:		parsed.Header,
				encryptedKey:	parsed.EncryptedKey.bytes(),
			},
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:156
		// _ = "end of CoverTab[186227]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:157
		_go_fuzz_dep_.CoverTab[186228]++
											obj.recipients = make([]recipientInfo, len(parsed.Recipients))
											for r := range parsed.Recipients {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:159
			_go_fuzz_dep_.CoverTab[186229]++
												encryptedKey, err := base64URLDecode(parsed.Recipients[r].EncryptedKey)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:161
				_go_fuzz_dep_.CoverTab[186232]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:162
				// _ = "end of CoverTab[186232]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:163
				_go_fuzz_dep_.CoverTab[186233]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:163
				// _ = "end of CoverTab[186233]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:163
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:163
			// _ = "end of CoverTab[186229]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:163
			_go_fuzz_dep_.CoverTab[186230]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:166
			if parsed.Recipients[r].Header != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:166
				_go_fuzz_dep_.CoverTab[186234]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:166
				return parsed.Recipients[r].Header.Nonce != ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:166
				// _ = "end of CoverTab[186234]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:166
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:166
				_go_fuzz_dep_.CoverTab[186235]++
													return nil, ErrUnprotectedNonce
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:167
				// _ = "end of CoverTab[186235]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:168
				_go_fuzz_dep_.CoverTab[186236]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:168
				// _ = "end of CoverTab[186236]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:168
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:168
			// _ = "end of CoverTab[186230]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:168
			_go_fuzz_dep_.CoverTab[186231]++

												obj.recipients[r].header = parsed.Recipients[r].Header
												obj.recipients[r].encryptedKey = encryptedKey
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:171
			// _ = "end of CoverTab[186231]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:172
		// _ = "end of CoverTab[186228]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:173
	// _ = "end of CoverTab[186214]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:173
	_go_fuzz_dep_.CoverTab[186215]++

										for _, recipient := range obj.recipients {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:175
		_go_fuzz_dep_.CoverTab[186237]++
											headers := obj.mergedHeaders(&recipient)
											if headers.Alg == "" || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:177
			_go_fuzz_dep_.CoverTab[186238]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:177
			return headers.Enc == ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:177
			// _ = "end of CoverTab[186238]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:177
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:177
			_go_fuzz_dep_.CoverTab[186239]++
												return nil, fmt.Errorf("square/go-jose: message is missing alg/enc headers")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:178
			// _ = "end of CoverTab[186239]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:179
			_go_fuzz_dep_.CoverTab[186240]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:179
			// _ = "end of CoverTab[186240]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:179
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:179
		// _ = "end of CoverTab[186237]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:180
	// _ = "end of CoverTab[186215]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:180
	_go_fuzz_dep_.CoverTab[186216]++

										obj.iv = parsed.Iv.bytes()
										obj.ciphertext = parsed.Ciphertext.bytes()
										obj.tag = parsed.Tag.bytes()
										obj.aad = parsed.Aad.bytes()

										return obj, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:187
	// _ = "end of CoverTab[186216]"
}

// parseEncryptedCompact parses a message in compact format.
func parseEncryptedCompact(input string) (*JsonWebEncryption, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:191
	_go_fuzz_dep_.CoverTab[186241]++
										parts := strings.Split(input, ".")
										if len(parts) != 5 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:193
		_go_fuzz_dep_.CoverTab[186248]++
											return nil, fmt.Errorf("square/go-jose: compact JWE format must have five parts")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:194
		// _ = "end of CoverTab[186248]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:195
		_go_fuzz_dep_.CoverTab[186249]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:195
		// _ = "end of CoverTab[186249]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:195
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:195
	// _ = "end of CoverTab[186241]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:195
	_go_fuzz_dep_.CoverTab[186242]++

										rawProtected, err := base64URLDecode(parts[0])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:198
		_go_fuzz_dep_.CoverTab[186250]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:199
		// _ = "end of CoverTab[186250]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:200
		_go_fuzz_dep_.CoverTab[186251]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:200
		// _ = "end of CoverTab[186251]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:200
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:200
	// _ = "end of CoverTab[186242]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:200
	_go_fuzz_dep_.CoverTab[186243]++

										encryptedKey, err := base64URLDecode(parts[1])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:203
		_go_fuzz_dep_.CoverTab[186252]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:204
		// _ = "end of CoverTab[186252]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:205
		_go_fuzz_dep_.CoverTab[186253]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:205
		// _ = "end of CoverTab[186253]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:205
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:205
	// _ = "end of CoverTab[186243]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:205
	_go_fuzz_dep_.CoverTab[186244]++

										iv, err := base64URLDecode(parts[2])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:208
		_go_fuzz_dep_.CoverTab[186254]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:209
		// _ = "end of CoverTab[186254]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:210
		_go_fuzz_dep_.CoverTab[186255]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:210
		// _ = "end of CoverTab[186255]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:210
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:210
	// _ = "end of CoverTab[186244]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:210
	_go_fuzz_dep_.CoverTab[186245]++

										ciphertext, err := base64URLDecode(parts[3])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:213
		_go_fuzz_dep_.CoverTab[186256]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:214
		// _ = "end of CoverTab[186256]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:215
		_go_fuzz_dep_.CoverTab[186257]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:215
		// _ = "end of CoverTab[186257]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:215
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:215
	// _ = "end of CoverTab[186245]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:215
	_go_fuzz_dep_.CoverTab[186246]++

										tag, err := base64URLDecode(parts[4])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:218
		_go_fuzz_dep_.CoverTab[186258]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:219
		// _ = "end of CoverTab[186258]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:220
		_go_fuzz_dep_.CoverTab[186259]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:220
		// _ = "end of CoverTab[186259]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:220
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:220
	// _ = "end of CoverTab[186246]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:220
	_go_fuzz_dep_.CoverTab[186247]++

										raw := &rawJsonWebEncryption{
		Protected:	newBuffer(rawProtected),
		EncryptedKey:	newBuffer(encryptedKey),
		Iv:		newBuffer(iv),
		Ciphertext:	newBuffer(ciphertext),
		Tag:		newBuffer(tag),
	}

										return raw.sanitized()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:230
	// _ = "end of CoverTab[186247]"
}

// CompactSerialize serializes an object using the compact serialization format.
func (obj JsonWebEncryption) CompactSerialize() (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:234
	_go_fuzz_dep_.CoverTab[186260]++
										if len(obj.recipients) != 1 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:235
		_go_fuzz_dep_.CoverTab[186262]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:235
		return obj.unprotected != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:235
		// _ = "end of CoverTab[186262]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:235
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:235
		_go_fuzz_dep_.CoverTab[186263]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:235
		return obj.protected == nil
											// _ = "end of CoverTab[186263]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:236
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:236
		_go_fuzz_dep_.CoverTab[186264]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:236
		return obj.recipients[0].header != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:236
		// _ = "end of CoverTab[186264]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:236
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:236
		_go_fuzz_dep_.CoverTab[186265]++
											return "", ErrNotSupported
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:237
		// _ = "end of CoverTab[186265]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:238
		_go_fuzz_dep_.CoverTab[186266]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:238
		// _ = "end of CoverTab[186266]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:238
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:238
	// _ = "end of CoverTab[186260]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:238
	_go_fuzz_dep_.CoverTab[186261]++

										serializedProtected := mustSerializeJSON(obj.protected)

										return fmt.Sprintf(
		"%s.%s.%s.%s.%s",
		base64URLEncode(serializedProtected),
		base64URLEncode(obj.recipients[0].encryptedKey),
		base64URLEncode(obj.iv),
		base64URLEncode(obj.ciphertext),
		base64URLEncode(obj.tag)), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:248
	// _ = "end of CoverTab[186261]"
}

// FullSerialize serializes an object using the full JSON serialization format.
func (obj JsonWebEncryption) FullSerialize() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:252
	_go_fuzz_dep_.CoverTab[186267]++
										raw := rawJsonWebEncryption{
		Unprotected:	obj.unprotected,
		Iv:		newBuffer(obj.iv),
		Ciphertext:	newBuffer(obj.ciphertext),
		EncryptedKey:	newBuffer(obj.recipients[0].encryptedKey),
		Tag:		newBuffer(obj.tag),
		Aad:		newBuffer(obj.aad),
		Recipients:	[]rawRecipientInfo{},
	}

	if len(obj.recipients) > 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:263
		_go_fuzz_dep_.CoverTab[186270]++
											for _, recipient := range obj.recipients {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:264
			_go_fuzz_dep_.CoverTab[186271]++
												info := rawRecipientInfo{
				Header:		recipient.header,
				EncryptedKey:	base64URLEncode(recipient.encryptedKey),
			}
												raw.Recipients = append(raw.Recipients, info)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:269
			// _ = "end of CoverTab[186271]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:270
		// _ = "end of CoverTab[186270]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:271
		_go_fuzz_dep_.CoverTab[186272]++

											raw.Header = obj.recipients[0].header
											raw.EncryptedKey = newBuffer(obj.recipients[0].encryptedKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:274
		// _ = "end of CoverTab[186272]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:275
	// _ = "end of CoverTab[186267]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:275
	_go_fuzz_dep_.CoverTab[186268]++

										if obj.protected != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:277
		_go_fuzz_dep_.CoverTab[186273]++
											raw.Protected = newBuffer(mustSerializeJSON(obj.protected))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:278
		// _ = "end of CoverTab[186273]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:279
		_go_fuzz_dep_.CoverTab[186274]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:279
		// _ = "end of CoverTab[186274]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:279
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:279
	// _ = "end of CoverTab[186268]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:279
	_go_fuzz_dep_.CoverTab[186269]++

										return string(mustSerializeJSON(raw))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:281
	// _ = "end of CoverTab[186269]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:282
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/jwe.go:282
var _ = _go_fuzz_dep_.CoverTab
