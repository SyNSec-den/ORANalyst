//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:1
package oidc

//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:1
)

import (
	"errors"

	"golang.org/x/oauth2"
)

// Nonce returns an auth code option which requires the ID Token created by the
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:9
// OpenID Connect provider to contain the specified nonce.
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:11
func Nonce(nonce string) oauth2.AuthCodeOption {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:11
	_go_fuzz_dep_.CoverTab[186836]++
														return oauth2.SetAuthURLParam("nonce", nonce)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:12
	// _ = "end of CoverTab[186836]"
}

// NonceSource represents a source which can verify a nonce is valid and has not
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:15
// been claimed before.
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:17
type NonceSource interface {
	ClaimNonce(nonce string) error
}

// VerifyNonce ensures that the ID Token contains a nonce which can be claimed by the nonce source.
func VerifyNonce(source NonceSource) VerificationOption {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:22
	_go_fuzz_dep_.CoverTab[186837]++
														return nonceVerifier{source}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:23
	// _ = "end of CoverTab[186837]"
}

type nonceVerifier struct {
	nonceSource NonceSource
}

func (n nonceVerifier) verifyIDToken(token *IDToken) error {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:30
	_go_fuzz_dep_.CoverTab[186838]++
														if token.Nonce == "" {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:31
		_go_fuzz_dep_.CoverTab[186840]++
															return errors.New("oidc: no nonce present in ID Token")
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:32
		// _ = "end of CoverTab[186840]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:33
		_go_fuzz_dep_.CoverTab[186841]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:33
		// _ = "end of CoverTab[186841]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:33
	// _ = "end of CoverTab[186838]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:33
	_go_fuzz_dep_.CoverTab[186839]++
														return n.nonceSource.ClaimNonce(token.Nonce)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:34
	// _ = "end of CoverTab[186839]"
}

//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:35
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/nonce.go:35
var _ = _go_fuzz_dep_.CoverTab
