//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:17
)

import (
	"encoding/base64"
	"fmt"
	"strings"

	"gopkg.in/square/go-jose.v2/json"
)

// rawJSONWebEncryption represents a raw JWE JSON object. Used for parsing/serializing.
type rawJSONWebEncryption struct {
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

// JSONWebEncryption represents an encrypted JWE object after parsing.
type JSONWebEncryption struct {
	Header				Header
	protected, unprotected		*rawHeader
	recipients			[]recipientInfo
	aad, iv, ciphertext, tag	[]byte
	original			*rawJSONWebEncryption
}

// recipientInfo represents a raw JWE Per-Recipient header JSON object after parsing.
type recipientInfo struct {
	header		*rawHeader
	encryptedKey	[]byte
}

// GetAuthData retrieves the (optional) authenticated data attached to the object.
func (obj JSONWebEncryption) GetAuthData() []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:62
	_go_fuzz_dep_.CoverTab[189341]++
										if obj.aad != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:63
		_go_fuzz_dep_.CoverTab[189343]++
											out := make([]byte, len(obj.aad))
											copy(out, obj.aad)
											return out
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:66
		// _ = "end of CoverTab[189343]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:67
		_go_fuzz_dep_.CoverTab[189344]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:67
		// _ = "end of CoverTab[189344]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:67
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:67
	// _ = "end of CoverTab[189341]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:67
	_go_fuzz_dep_.CoverTab[189342]++

										return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:69
	// _ = "end of CoverTab[189342]"
}

// Get the merged header values
func (obj JSONWebEncryption) mergedHeaders(recipient *recipientInfo) rawHeader {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:73
	_go_fuzz_dep_.CoverTab[189345]++
										out := rawHeader{}
										out.merge(obj.protected)
										out.merge(obj.unprotected)

										if recipient != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:78
		_go_fuzz_dep_.CoverTab[189347]++
											out.merge(recipient.header)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:79
		// _ = "end of CoverTab[189347]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:80
		_go_fuzz_dep_.CoverTab[189348]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:80
		// _ = "end of CoverTab[189348]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:80
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:80
	// _ = "end of CoverTab[189345]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:80
	_go_fuzz_dep_.CoverTab[189346]++

										return out
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:82
	// _ = "end of CoverTab[189346]"
}

// Get the additional authenticated data from a JWE object.
func (obj JSONWebEncryption) computeAuthData() []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:86
	_go_fuzz_dep_.CoverTab[189349]++
										var protected string

										if obj.original != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:89
		_go_fuzz_dep_.CoverTab[189352]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:89
		return obj.original.Protected != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:89
		// _ = "end of CoverTab[189352]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:89
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:89
		_go_fuzz_dep_.CoverTab[189353]++
											protected = obj.original.Protected.base64()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:90
		// _ = "end of CoverTab[189353]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:91
		_go_fuzz_dep_.CoverTab[189354]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:91
		if obj.protected != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:91
			_go_fuzz_dep_.CoverTab[189355]++
												protected = base64.RawURLEncoding.EncodeToString(mustSerializeJSON((obj.protected)))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:92
			// _ = "end of CoverTab[189355]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:93
			_go_fuzz_dep_.CoverTab[189356]++
												protected = ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:94
			// _ = "end of CoverTab[189356]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:95
		// _ = "end of CoverTab[189354]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:95
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:95
	// _ = "end of CoverTab[189349]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:95
	_go_fuzz_dep_.CoverTab[189350]++

										output := []byte(protected)
										if obj.aad != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:98
		_go_fuzz_dep_.CoverTab[189357]++
											output = append(output, '.')
											output = append(output, []byte(base64.RawURLEncoding.EncodeToString(obj.aad))...)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:100
		// _ = "end of CoverTab[189357]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:101
		_go_fuzz_dep_.CoverTab[189358]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:101
		// _ = "end of CoverTab[189358]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:101
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:101
	// _ = "end of CoverTab[189350]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:101
	_go_fuzz_dep_.CoverTab[189351]++

										return output
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:103
	// _ = "end of CoverTab[189351]"
}

// ParseEncrypted parses an encrypted message in compact or full serialization format.
func ParseEncrypted(input string) (*JSONWebEncryption, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:107
	_go_fuzz_dep_.CoverTab[189359]++
										input = stripWhitespace(input)
										if strings.HasPrefix(input, "{") {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:109
		_go_fuzz_dep_.CoverTab[189361]++
											return parseEncryptedFull(input)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:110
		// _ = "end of CoverTab[189361]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:111
		_go_fuzz_dep_.CoverTab[189362]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:111
		// _ = "end of CoverTab[189362]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:111
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:111
	// _ = "end of CoverTab[189359]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:111
	_go_fuzz_dep_.CoverTab[189360]++

										return parseEncryptedCompact(input)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:113
	// _ = "end of CoverTab[189360]"
}

// parseEncryptedFull parses a message in compact format.
func parseEncryptedFull(input string) (*JSONWebEncryption, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:117
	_go_fuzz_dep_.CoverTab[189363]++
										var parsed rawJSONWebEncryption
										err := json.Unmarshal([]byte(input), &parsed)
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:120
		_go_fuzz_dep_.CoverTab[189365]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:121
		// _ = "end of CoverTab[189365]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:122
		_go_fuzz_dep_.CoverTab[189366]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:122
		// _ = "end of CoverTab[189366]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:122
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:122
	// _ = "end of CoverTab[189363]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:122
	_go_fuzz_dep_.CoverTab[189364]++

										return parsed.sanitized()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:124
	// _ = "end of CoverTab[189364]"
}

// sanitized produces a cleaned-up JWE object from the raw JSON.
func (parsed *rawJSONWebEncryption) sanitized() (*JSONWebEncryption, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:128
	_go_fuzz_dep_.CoverTab[189367]++
										obj := &JSONWebEncryption{
		original:	parsed,
		unprotected:	parsed.Unprotected,
	}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:135
	if parsed.Unprotected != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:135
		_go_fuzz_dep_.CoverTab[189374]++
											if nonce := parsed.Unprotected.getNonce(); nonce != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:136
			_go_fuzz_dep_.CoverTab[189375]++
												return nil, ErrUnprotectedNonce
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:137
			// _ = "end of CoverTab[189375]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:138
			_go_fuzz_dep_.CoverTab[189376]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:138
			// _ = "end of CoverTab[189376]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:138
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:138
		// _ = "end of CoverTab[189374]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:139
		_go_fuzz_dep_.CoverTab[189377]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:139
		// _ = "end of CoverTab[189377]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:139
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:139
	// _ = "end of CoverTab[189367]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:139
	_go_fuzz_dep_.CoverTab[189368]++
										if parsed.Header != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:140
		_go_fuzz_dep_.CoverTab[189378]++
											if nonce := parsed.Header.getNonce(); nonce != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:141
			_go_fuzz_dep_.CoverTab[189379]++
												return nil, ErrUnprotectedNonce
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:142
			// _ = "end of CoverTab[189379]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:143
			_go_fuzz_dep_.CoverTab[189380]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:143
			// _ = "end of CoverTab[189380]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:143
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:143
		// _ = "end of CoverTab[189378]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:144
		_go_fuzz_dep_.CoverTab[189381]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:144
		// _ = "end of CoverTab[189381]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:144
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:144
	// _ = "end of CoverTab[189368]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:144
	_go_fuzz_dep_.CoverTab[189369]++

										if parsed.Protected != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:146
		_go_fuzz_dep_.CoverTab[189382]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:146
		return len(parsed.Protected.bytes()) > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:146
		// _ = "end of CoverTab[189382]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:146
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:146
		_go_fuzz_dep_.CoverTab[189383]++
											err := json.Unmarshal(parsed.Protected.bytes(), &obj.protected)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:148
			_go_fuzz_dep_.CoverTab[189384]++
												return nil, fmt.Errorf("square/go-jose: invalid protected header: %s, %s", err, parsed.Protected.base64())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:149
			// _ = "end of CoverTab[189384]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:150
			_go_fuzz_dep_.CoverTab[189385]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:150
			// _ = "end of CoverTab[189385]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:150
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:150
		// _ = "end of CoverTab[189383]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:151
		_go_fuzz_dep_.CoverTab[189386]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:151
		// _ = "end of CoverTab[189386]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:151
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:151
	// _ = "end of CoverTab[189369]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:151
	_go_fuzz_dep_.CoverTab[189370]++

	// Note: this must be called _after_ we parse the protected header,
	// otherwise fields from the protected header will not get picked up.
	var err error
	mergedHeaders := obj.mergedHeaders(nil)
	obj.Header, err = mergedHeaders.sanitized()
	if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:158
		_go_fuzz_dep_.CoverTab[189387]++
											return nil, fmt.Errorf("square/go-jose: cannot sanitize merged headers: %v (%v)", err, mergedHeaders)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:159
		// _ = "end of CoverTab[189387]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:160
		_go_fuzz_dep_.CoverTab[189388]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:160
		// _ = "end of CoverTab[189388]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:160
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:160
	// _ = "end of CoverTab[189370]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:160
	_go_fuzz_dep_.CoverTab[189371]++

										if len(parsed.Recipients) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:162
		_go_fuzz_dep_.CoverTab[189389]++
											obj.recipients = []recipientInfo{
			{
				header:		parsed.Header,
				encryptedKey:	parsed.EncryptedKey.bytes(),
			},
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:168
		// _ = "end of CoverTab[189389]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:169
		_go_fuzz_dep_.CoverTab[189390]++
											obj.recipients = make([]recipientInfo, len(parsed.Recipients))
											for r := range parsed.Recipients {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:171
			_go_fuzz_dep_.CoverTab[189391]++
												encryptedKey, err := base64.RawURLEncoding.DecodeString(parsed.Recipients[r].EncryptedKey)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:173
				_go_fuzz_dep_.CoverTab[189394]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:174
				// _ = "end of CoverTab[189394]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:175
				_go_fuzz_dep_.CoverTab[189395]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:175
				// _ = "end of CoverTab[189395]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:175
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:175
			// _ = "end of CoverTab[189391]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:175
			_go_fuzz_dep_.CoverTab[189392]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:178
			if parsed.Recipients[r].Header != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:178
				_go_fuzz_dep_.CoverTab[189396]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:178
				return parsed.Recipients[r].Header.getNonce() != ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:178
				// _ = "end of CoverTab[189396]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:178
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:178
				_go_fuzz_dep_.CoverTab[189397]++
													return nil, ErrUnprotectedNonce
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:179
				// _ = "end of CoverTab[189397]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:180
				_go_fuzz_dep_.CoverTab[189398]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:180
				// _ = "end of CoverTab[189398]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:180
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:180
			// _ = "end of CoverTab[189392]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:180
			_go_fuzz_dep_.CoverTab[189393]++

												obj.recipients[r].header = parsed.Recipients[r].Header
												obj.recipients[r].encryptedKey = encryptedKey
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:183
			// _ = "end of CoverTab[189393]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:184
		// _ = "end of CoverTab[189390]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:185
	// _ = "end of CoverTab[189371]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:185
	_go_fuzz_dep_.CoverTab[189372]++

										for _, recipient := range obj.recipients {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:187
		_go_fuzz_dep_.CoverTab[189399]++
											headers := obj.mergedHeaders(&recipient)
											if headers.getAlgorithm() == "" || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:189
			_go_fuzz_dep_.CoverTab[189400]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:189
			return headers.getEncryption() == ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:189
			// _ = "end of CoverTab[189400]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:189
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:189
			_go_fuzz_dep_.CoverTab[189401]++
												return nil, fmt.Errorf("square/go-jose: message is missing alg/enc headers")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:190
			// _ = "end of CoverTab[189401]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:191
			_go_fuzz_dep_.CoverTab[189402]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:191
			// _ = "end of CoverTab[189402]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:191
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:191
		// _ = "end of CoverTab[189399]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:192
	// _ = "end of CoverTab[189372]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:192
	_go_fuzz_dep_.CoverTab[189373]++

										obj.iv = parsed.Iv.bytes()
										obj.ciphertext = parsed.Ciphertext.bytes()
										obj.tag = parsed.Tag.bytes()
										obj.aad = parsed.Aad.bytes()

										return obj, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:199
	// _ = "end of CoverTab[189373]"
}

// parseEncryptedCompact parses a message in compact format.
func parseEncryptedCompact(input string) (*JSONWebEncryption, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:203
	_go_fuzz_dep_.CoverTab[189403]++
										parts := strings.Split(input, ".")
										if len(parts) != 5 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:205
		_go_fuzz_dep_.CoverTab[189410]++
											return nil, fmt.Errorf("square/go-jose: compact JWE format must have five parts")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:206
		// _ = "end of CoverTab[189410]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:207
		_go_fuzz_dep_.CoverTab[189411]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:207
		// _ = "end of CoverTab[189411]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:207
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:207
	// _ = "end of CoverTab[189403]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:207
	_go_fuzz_dep_.CoverTab[189404]++

										rawProtected, err := base64.RawURLEncoding.DecodeString(parts[0])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:210
		_go_fuzz_dep_.CoverTab[189412]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:211
		// _ = "end of CoverTab[189412]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:212
		_go_fuzz_dep_.CoverTab[189413]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:212
		// _ = "end of CoverTab[189413]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:212
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:212
	// _ = "end of CoverTab[189404]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:212
	_go_fuzz_dep_.CoverTab[189405]++

										encryptedKey, err := base64.RawURLEncoding.DecodeString(parts[1])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:215
		_go_fuzz_dep_.CoverTab[189414]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:216
		// _ = "end of CoverTab[189414]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:217
		_go_fuzz_dep_.CoverTab[189415]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:217
		// _ = "end of CoverTab[189415]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:217
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:217
	// _ = "end of CoverTab[189405]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:217
	_go_fuzz_dep_.CoverTab[189406]++

										iv, err := base64.RawURLEncoding.DecodeString(parts[2])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:220
		_go_fuzz_dep_.CoverTab[189416]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:221
		// _ = "end of CoverTab[189416]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:222
		_go_fuzz_dep_.CoverTab[189417]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:222
		// _ = "end of CoverTab[189417]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:222
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:222
	// _ = "end of CoverTab[189406]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:222
	_go_fuzz_dep_.CoverTab[189407]++

										ciphertext, err := base64.RawURLEncoding.DecodeString(parts[3])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:225
		_go_fuzz_dep_.CoverTab[189418]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:226
		// _ = "end of CoverTab[189418]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:227
		_go_fuzz_dep_.CoverTab[189419]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:227
		// _ = "end of CoverTab[189419]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:227
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:227
	// _ = "end of CoverTab[189407]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:227
	_go_fuzz_dep_.CoverTab[189408]++

										tag, err := base64.RawURLEncoding.DecodeString(parts[4])
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:230
		_go_fuzz_dep_.CoverTab[189420]++
											return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:231
		// _ = "end of CoverTab[189420]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:232
		_go_fuzz_dep_.CoverTab[189421]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:232
		// _ = "end of CoverTab[189421]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:232
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:232
	// _ = "end of CoverTab[189408]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:232
	_go_fuzz_dep_.CoverTab[189409]++

										raw := &rawJSONWebEncryption{
		Protected:	newBuffer(rawProtected),
		EncryptedKey:	newBuffer(encryptedKey),
		Iv:		newBuffer(iv),
		Ciphertext:	newBuffer(ciphertext),
		Tag:		newBuffer(tag),
	}

										return raw.sanitized()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:242
	// _ = "end of CoverTab[189409]"
}

// CompactSerialize serializes an object using the compact serialization format.
func (obj JSONWebEncryption) CompactSerialize() (string, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:246
	_go_fuzz_dep_.CoverTab[189422]++
										if len(obj.recipients) != 1 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:247
		_go_fuzz_dep_.CoverTab[189424]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:247
		return obj.unprotected != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:247
		// _ = "end of CoverTab[189424]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:247
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:247
		_go_fuzz_dep_.CoverTab[189425]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:247
		return obj.protected == nil
											// _ = "end of CoverTab[189425]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:248
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:248
		_go_fuzz_dep_.CoverTab[189426]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:248
		return obj.recipients[0].header != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:248
		// _ = "end of CoverTab[189426]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:248
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:248
		_go_fuzz_dep_.CoverTab[189427]++
											return "", ErrNotSupported
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:249
		// _ = "end of CoverTab[189427]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:250
		_go_fuzz_dep_.CoverTab[189428]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:250
		// _ = "end of CoverTab[189428]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:250
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:250
	// _ = "end of CoverTab[189422]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:250
	_go_fuzz_dep_.CoverTab[189423]++

										serializedProtected := mustSerializeJSON(obj.protected)

										return fmt.Sprintf(
		"%s.%s.%s.%s.%s",
		base64.RawURLEncoding.EncodeToString(serializedProtected),
		base64.RawURLEncoding.EncodeToString(obj.recipients[0].encryptedKey),
		base64.RawURLEncoding.EncodeToString(obj.iv),
		base64.RawURLEncoding.EncodeToString(obj.ciphertext),
		base64.RawURLEncoding.EncodeToString(obj.tag)), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:260
	// _ = "end of CoverTab[189423]"
}

// FullSerialize serializes an object using the full JSON serialization format.
func (obj JSONWebEncryption) FullSerialize() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:264
	_go_fuzz_dep_.CoverTab[189429]++
										raw := rawJSONWebEncryption{
		Unprotected:	obj.unprotected,
		Iv:		newBuffer(obj.iv),
		Ciphertext:	newBuffer(obj.ciphertext),
		EncryptedKey:	newBuffer(obj.recipients[0].encryptedKey),
		Tag:		newBuffer(obj.tag),
		Aad:		newBuffer(obj.aad),
		Recipients:	[]rawRecipientInfo{},
	}

	if len(obj.recipients) > 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:275
		_go_fuzz_dep_.CoverTab[189432]++
											for _, recipient := range obj.recipients {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:276
			_go_fuzz_dep_.CoverTab[189433]++
												info := rawRecipientInfo{
				Header:		recipient.header,
				EncryptedKey:	base64.RawURLEncoding.EncodeToString(recipient.encryptedKey),
			}
												raw.Recipients = append(raw.Recipients, info)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:281
			// _ = "end of CoverTab[189433]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:282
		// _ = "end of CoverTab[189432]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:283
		_go_fuzz_dep_.CoverTab[189434]++

											raw.Header = obj.recipients[0].header
											raw.EncryptedKey = newBuffer(obj.recipients[0].encryptedKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:286
		// _ = "end of CoverTab[189434]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:287
	// _ = "end of CoverTab[189429]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:287
	_go_fuzz_dep_.CoverTab[189430]++

										if obj.protected != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:289
		_go_fuzz_dep_.CoverTab[189435]++
											raw.Protected = newBuffer(mustSerializeJSON(obj.protected))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:290
		// _ = "end of CoverTab[189435]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:291
		_go_fuzz_dep_.CoverTab[189436]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:291
		// _ = "end of CoverTab[189436]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:291
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:291
	// _ = "end of CoverTab[189430]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:291
	_go_fuzz_dep_.CoverTab[189431]++

										return string(mustSerializeJSON(raw))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:293
	// _ = "end of CoverTab[189431]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:294
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/jwe.go:294
var _ = _go_fuzz_dep_.CoverTab
