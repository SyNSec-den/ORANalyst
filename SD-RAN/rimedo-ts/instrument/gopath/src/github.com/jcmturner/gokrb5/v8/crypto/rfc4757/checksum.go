//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:1
package rfc4757

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:1
)

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"io"
)

// Checksum returns a hash of the data in accordance with RFC 4757
func Checksum(key []byte, usage uint32, data []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:11
	_go_fuzz_dep_.CoverTab[85799]++

													s := append([]byte(`signaturekey`), byte(0x00))
													mac := hmac.New(md5.New, key)
													mac.Write(s)
													Ksign := mac.Sum(nil)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:19
	tb := UsageToMSMsgType(usage)
	p := append(tb, data...)
	h := md5.New()
	rb := bytes.NewReader(p)
	_, err := io.Copy(h, rb)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:24
		_go_fuzz_dep_.CoverTab[85801]++
														return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:25
		// _ = "end of CoverTab[85801]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:26
		_go_fuzz_dep_.CoverTab[85802]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:26
		// _ = "end of CoverTab[85802]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:26
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:26
	// _ = "end of CoverTab[85799]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:26
	_go_fuzz_dep_.CoverTab[85800]++
													tmp := h.Sum(nil)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:30
	mac = hmac.New(md5.New, Ksign)
													mac.Write(tmp)
													return mac.Sum(nil), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:32
	// _ = "end of CoverTab[85800]"
}

// HMAC returns a keyed MD5 checksum of the data
func HMAC(key []byte, data []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:36
	_go_fuzz_dep_.CoverTab[85803]++
													mac := hmac.New(md5.New, key)
													mac.Write(data)
													return mac.Sum(nil)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:39
	// _ = "end of CoverTab[85803]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:40
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/checksum.go:40
var _ = _go_fuzz_dep_.CoverTab
