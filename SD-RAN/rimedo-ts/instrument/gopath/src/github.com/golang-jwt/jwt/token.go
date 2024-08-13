//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:1
package jwt

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:1
)

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"
)

// TimeFunc provides the current time when parsing token to validate "exp" claim (expiration time).
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:10
// You can override it to use another time value.  This is useful for testing or if your
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:10
// server uses a different time zone than your tokens.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:13
var TimeFunc = time.Now

// Parse methods use this callback function to supply
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:15
// the key for verification.  The function receives the parsed,
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:15
// but unverified Token.  This allows you to use properties in the
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:15
// Header of the token (such as `kid`) to identify which key to use.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:19
type Keyfunc func(*Token) (interface{}, error)

// A JWT Token.  Different fields will be used depending on whether you're
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:21
// creating or parsing/verifying a token.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:23
type Token struct {
	Raw		string			// The raw token.  Populated when you Parse a token
	Method		SigningMethod		// The signing method used or to be used
	Header		map[string]interface{}	// The first segment of the token
	Claims		Claims			// The second segment of the token
	Signature	string			// The third segment of the token.  Populated when you Parse a token
	Valid		bool			// Is the token valid?  Populated when you Parse/Verify a token
}

// Create a new Token.  Takes a signing method
func New(method SigningMethod) *Token {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:33
	_go_fuzz_dep_.CoverTab[187317]++
												return NewWithClaims(method, MapClaims{})
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:34
	// _ = "end of CoverTab[187317]"
}

func NewWithClaims(method SigningMethod, claims Claims) *Token {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:37
	_go_fuzz_dep_.CoverTab[187318]++
												return &Token{
		Header: map[string]interface{}{
			"typ":	"JWT",
			"alg":	method.Alg(),
		},
		Claims:	claims,
		Method:	method,
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:45
	// _ = "end of CoverTab[187318]"
}

// Get the complete, signed token
func (t *Token) SignedString(key interface{}) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:49
	_go_fuzz_dep_.CoverTab[187319]++
												var sig, sstr string
												var err error
												if sstr, err = t.SigningString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:52
		_go_fuzz_dep_.CoverTab[187322]++
													return "", err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:53
		// _ = "end of CoverTab[187322]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:54
		_go_fuzz_dep_.CoverTab[187323]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:54
		// _ = "end of CoverTab[187323]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:54
	// _ = "end of CoverTab[187319]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:54
	_go_fuzz_dep_.CoverTab[187320]++
												if sig, err = t.Method.Sign(sstr, key); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:55
		_go_fuzz_dep_.CoverTab[187324]++
													return "", err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:56
		// _ = "end of CoverTab[187324]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:57
		_go_fuzz_dep_.CoverTab[187325]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:57
		// _ = "end of CoverTab[187325]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:57
	// _ = "end of CoverTab[187320]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:57
	_go_fuzz_dep_.CoverTab[187321]++
												return strings.Join([]string{sstr, sig}, "."), nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:58
	// _ = "end of CoverTab[187321]"
}

// Generate the signing string.  This is the
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:61
// most expensive part of the whole deal.  Unless you
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:61
// need this for something special, just go straight for
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:61
// the SignedString.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:65
func (t *Token) SigningString() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:65
	_go_fuzz_dep_.CoverTab[187326]++
												var err error
												parts := make([]string, 2)
												for i := range parts {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:68
		_go_fuzz_dep_.CoverTab[187328]++
													var jsonValue []byte
													if i == 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:70
			_go_fuzz_dep_.CoverTab[187330]++
														if jsonValue, err = json.Marshal(t.Header); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:71
				_go_fuzz_dep_.CoverTab[187331]++
															return "", err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:72
				// _ = "end of CoverTab[187331]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:73
				_go_fuzz_dep_.CoverTab[187332]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:73
				// _ = "end of CoverTab[187332]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:73
			}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:73
			// _ = "end of CoverTab[187330]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:74
			_go_fuzz_dep_.CoverTab[187333]++
														if jsonValue, err = json.Marshal(t.Claims); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:75
				_go_fuzz_dep_.CoverTab[187334]++
															return "", err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:76
				// _ = "end of CoverTab[187334]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:77
				_go_fuzz_dep_.CoverTab[187335]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:77
				// _ = "end of CoverTab[187335]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:77
			}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:77
			// _ = "end of CoverTab[187333]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:78
		// _ = "end of CoverTab[187328]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:78
		_go_fuzz_dep_.CoverTab[187329]++

													parts[i] = EncodeSegment(jsonValue)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:80
		// _ = "end of CoverTab[187329]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:81
	// _ = "end of CoverTab[187326]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:81
	_go_fuzz_dep_.CoverTab[187327]++
												return strings.Join(parts, "."), nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:82
	// _ = "end of CoverTab[187327]"
}

// Parse, validate, and return a token.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:85
// keyFunc will receive the parsed token and should return the key for validating.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:85
// If everything is kosher, err will be nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:88
func Parse(tokenString string, keyFunc Keyfunc) (*Token, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:88
	_go_fuzz_dep_.CoverTab[187336]++
												return new(Parser).Parse(tokenString, keyFunc)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:89
	// _ = "end of CoverTab[187336]"
}

func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:92
	_go_fuzz_dep_.CoverTab[187337]++
												return new(Parser).ParseWithClaims(tokenString, claims, keyFunc)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:93
	// _ = "end of CoverTab[187337]"
}

// Encode JWT specific base64url encoding with padding stripped
func EncodeSegment(seg []byte) string {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:97
	_go_fuzz_dep_.CoverTab[187338]++
												return base64.RawURLEncoding.EncodeToString(seg)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:98
	// _ = "end of CoverTab[187338]"
}

// Decode JWT specific base64url encoding with padding stripped
func DecodeSegment(seg string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:102
	_go_fuzz_dep_.CoverTab[187339]++
												return base64.RawURLEncoding.DecodeString(seg)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:103
	// _ = "end of CoverTab[187339]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:104
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/token.go:104
var _ = _go_fuzz_dep_.CoverTab
