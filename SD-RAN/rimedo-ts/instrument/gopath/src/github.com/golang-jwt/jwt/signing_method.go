//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:1
package jwt

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:1
)

import (
	"sync"
)

var signingMethods = map[string]func() SigningMethod{}
var signingMethodLock = new(sync.RWMutex)

// Implement SigningMethod to add new methods for signing or verifying tokens.
type SigningMethod interface {
	Verify(signingString, signature string, key interface{}) error	// Returns nil if signature is valid
	Sign(signingString string, key interface{}) (string, error)	// Returns encoded signature or error
	Alg() string							// returns the alg identifier for this method (example: 'HS256')
}

// Register the "alg" name and a factory function for signing method.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:17
// This is typically done during init() in the method's implementation
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:19
func RegisterSigningMethod(alg string, f func() SigningMethod) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:19
	_go_fuzz_dep_.CoverTab[187312]++
													signingMethodLock.Lock()
													defer signingMethodLock.Unlock()

													signingMethods[alg] = f
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:23
	// _ = "end of CoverTab[187312]"
}

// Get a signing method from an "alg" string
func GetSigningMethod(alg string) (method SigningMethod) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:27
	_go_fuzz_dep_.CoverTab[187313]++
													signingMethodLock.RLock()
													defer signingMethodLock.RUnlock()

													if methodF, ok := signingMethods[alg]; ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:31
		_go_fuzz_dep_.CoverTab[187315]++
														method = methodF()
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:32
		// _ = "end of CoverTab[187315]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:33
		_go_fuzz_dep_.CoverTab[187316]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:33
		// _ = "end of CoverTab[187316]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:33
	// _ = "end of CoverTab[187313]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:33
	_go_fuzz_dep_.CoverTab[187314]++
													return
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:34
	// _ = "end of CoverTab[187314]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:35
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/signing_method.go:35
var _ = _go_fuzz_dep_.CoverTab
