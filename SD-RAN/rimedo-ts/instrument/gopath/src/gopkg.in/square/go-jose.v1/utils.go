//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:17
)

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// LoadPublicKey loads a public key from PEM/DER-encoded data.
func LoadPublicKey(data []byte) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:26
	_go_fuzz_dep_.CoverTab[186748]++
											input := data

											block, _ := pem.Decode(data)
											if block != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:30
		_go_fuzz_dep_.CoverTab[186752]++
												input = block.Bytes
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:31
		// _ = "end of CoverTab[186752]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:32
		_go_fuzz_dep_.CoverTab[186753]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:32
		// _ = "end of CoverTab[186753]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:32
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:32
	// _ = "end of CoverTab[186748]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:32
	_go_fuzz_dep_.CoverTab[186749]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:35
	pub, err0 := x509.ParsePKIXPublicKey(input)
	if err0 == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:36
		_go_fuzz_dep_.CoverTab[186754]++
												return pub, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:37
		// _ = "end of CoverTab[186754]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:38
		_go_fuzz_dep_.CoverTab[186755]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:38
		// _ = "end of CoverTab[186755]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:38
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:38
	// _ = "end of CoverTab[186749]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:38
	_go_fuzz_dep_.CoverTab[186750]++

											cert, err1 := x509.ParseCertificate(input)
											if err1 == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:41
		_go_fuzz_dep_.CoverTab[186756]++
												return cert.PublicKey, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:42
		// _ = "end of CoverTab[186756]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:43
		_go_fuzz_dep_.CoverTab[186757]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:43
		// _ = "end of CoverTab[186757]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:43
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:43
	// _ = "end of CoverTab[186750]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:43
	_go_fuzz_dep_.CoverTab[186751]++

											return nil, fmt.Errorf("square/go-jose: parse error, got '%s' and '%s'", err0, err1)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:45
	// _ = "end of CoverTab[186751]"
}

// LoadPrivateKey loads a private key from PEM/DER-encoded data.
func LoadPrivateKey(data []byte) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:49
	_go_fuzz_dep_.CoverTab[186758]++
											input := data

											block, _ := pem.Decode(data)
											if block != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:53
		_go_fuzz_dep_.CoverTab[186763]++
												input = block.Bytes
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:54
		// _ = "end of CoverTab[186763]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:55
		_go_fuzz_dep_.CoverTab[186764]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:55
		// _ = "end of CoverTab[186764]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:55
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:55
	// _ = "end of CoverTab[186758]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:55
	_go_fuzz_dep_.CoverTab[186759]++

											var priv interface{}
											priv, err0 := x509.ParsePKCS1PrivateKey(input)
											if err0 == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:59
		_go_fuzz_dep_.CoverTab[186765]++
												return priv, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:60
		// _ = "end of CoverTab[186765]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:61
		_go_fuzz_dep_.CoverTab[186766]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:61
		// _ = "end of CoverTab[186766]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:61
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:61
	// _ = "end of CoverTab[186759]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:61
	_go_fuzz_dep_.CoverTab[186760]++

											priv, err1 := x509.ParsePKCS8PrivateKey(input)
											if err1 == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:64
		_go_fuzz_dep_.CoverTab[186767]++
												return priv, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:65
		// _ = "end of CoverTab[186767]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:66
		_go_fuzz_dep_.CoverTab[186768]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:66
		// _ = "end of CoverTab[186768]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:66
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:66
	// _ = "end of CoverTab[186760]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:66
	_go_fuzz_dep_.CoverTab[186761]++

											priv, err2 := x509.ParseECPrivateKey(input)
											if err2 == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:69
		_go_fuzz_dep_.CoverTab[186769]++
												return priv, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:70
		// _ = "end of CoverTab[186769]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:71
		_go_fuzz_dep_.CoverTab[186770]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:71
		// _ = "end of CoverTab[186770]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:71
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:71
	// _ = "end of CoverTab[186761]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:71
	_go_fuzz_dep_.CoverTab[186762]++

											return nil, fmt.Errorf("square/go-jose: parse error, got '%s', '%s' and '%s'", err0, err1, err2)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:73
	// _ = "end of CoverTab[186762]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:74
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/utils.go:74
var _ = _go_fuzz_dep_.CoverTab
