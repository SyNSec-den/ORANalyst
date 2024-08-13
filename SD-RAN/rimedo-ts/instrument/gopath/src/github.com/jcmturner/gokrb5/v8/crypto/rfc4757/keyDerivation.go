//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:1
package rfc4757

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:1
)

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/md4"
)

// StringToKey returns a key derived from the string provided according to the definition in RFC 4757.
func StringToKey(secret string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:14
	_go_fuzz_dep_.CoverTab[85827]++
														b := make([]byte, len(secret)*2, len(secret)*2)
														for i, r := range secret {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:16
		_go_fuzz_dep_.CoverTab[85830]++
															u := fmt.Sprintf("%04x", r)
															c, err := hex.DecodeString(u)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:19
			_go_fuzz_dep_.CoverTab[85832]++
																return []byte{}, errors.New("character could not be encoded")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:20
			// _ = "end of CoverTab[85832]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:21
			_go_fuzz_dep_.CoverTab[85833]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:21
			// _ = "end of CoverTab[85833]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:21
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:21
		// _ = "end of CoverTab[85830]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:21
		_go_fuzz_dep_.CoverTab[85831]++

															b[2*i] = c[1]
															b[2*i+1] = c[0]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:24
		// _ = "end of CoverTab[85831]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:25
	// _ = "end of CoverTab[85827]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:25
	_go_fuzz_dep_.CoverTab[85828]++
														r := bytes.NewReader(b)
														h := md4.New()
														_, err := io.Copy(h, r)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:29
		_go_fuzz_dep_.CoverTab[85834]++
															return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:30
		// _ = "end of CoverTab[85834]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:31
		_go_fuzz_dep_.CoverTab[85835]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:31
		// _ = "end of CoverTab[85835]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:31
	// _ = "end of CoverTab[85828]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:31
	_go_fuzz_dep_.CoverTab[85829]++
														return h.Sum(nil), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:32
	// _ = "end of CoverTab[85829]"
}

func deriveKeys(key, checksum []byte, usage uint32, export bool) (k1, k2, k3 []byte) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:35
	_go_fuzz_dep_.CoverTab[85836]++
														k1 = key
														k2 = HMAC(k1, UsageToMSMsgType(usage))
														k3 = HMAC(k2, checksum)
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:39
	// _ = "end of CoverTab[85836]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:40
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/keyDerivation.go:40
var _ = _go_fuzz_dep_.CoverTab
