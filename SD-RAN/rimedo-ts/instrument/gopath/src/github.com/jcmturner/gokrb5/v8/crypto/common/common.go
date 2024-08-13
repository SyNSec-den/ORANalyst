//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:1
// Package common provides encryption methods common across encryption types
package common

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:2
)

import (
	"bytes"
	"crypto/hmac"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/jcmturner/gokrb5/v8/crypto/etype"
)

// ZeroPad pads bytes with zeros to nearest multiple of message size m.
func ZeroPad(b []byte, m int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:16
	_go_fuzz_dep_.CoverTab[85474]++
													if m <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:17
		_go_fuzz_dep_.CoverTab[85478]++
														return nil, errors.New("Invalid message block size when padding")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:18
		// _ = "end of CoverTab[85478]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:19
		_go_fuzz_dep_.CoverTab[85479]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:19
		// _ = "end of CoverTab[85479]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:19
	// _ = "end of CoverTab[85474]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:19
	_go_fuzz_dep_.CoverTab[85475]++
													if b == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:20
		_go_fuzz_dep_.CoverTab[85480]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:20
		return len(b) == 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:20
		// _ = "end of CoverTab[85480]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:20
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:20
		_go_fuzz_dep_.CoverTab[85481]++
														return nil, errors.New("Data not valid to pad: Zero size")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:21
		// _ = "end of CoverTab[85481]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:22
		_go_fuzz_dep_.CoverTab[85482]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:22
		// _ = "end of CoverTab[85482]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:22
	// _ = "end of CoverTab[85475]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:22
	_go_fuzz_dep_.CoverTab[85476]++
													if l := len(b) % m; l != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:23
		_go_fuzz_dep_.CoverTab[85483]++
														n := m - l
														z := make([]byte, n)
														b = append(b, z...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:26
		// _ = "end of CoverTab[85483]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:27
		_go_fuzz_dep_.CoverTab[85484]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:27
		// _ = "end of CoverTab[85484]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:27
	// _ = "end of CoverTab[85476]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:27
	_go_fuzz_dep_.CoverTab[85477]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:28
	// _ = "end of CoverTab[85477]"
}

// PKCS7Pad pads bytes according to RFC 2315 to nearest multiple of message size m.
func PKCS7Pad(b []byte, m int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:32
	_go_fuzz_dep_.CoverTab[85485]++
													if m <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:33
		_go_fuzz_dep_.CoverTab[85488]++
														return nil, errors.New("Invalid message block size when padding")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:34
		// _ = "end of CoverTab[85488]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:35
		_go_fuzz_dep_.CoverTab[85489]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:35
		// _ = "end of CoverTab[85489]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:35
	// _ = "end of CoverTab[85485]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:35
	_go_fuzz_dep_.CoverTab[85486]++
													if b == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:36
		_go_fuzz_dep_.CoverTab[85490]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:36
		return len(b) == 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:36
		// _ = "end of CoverTab[85490]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:36
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:36
		_go_fuzz_dep_.CoverTab[85491]++
														return nil, errors.New("Data not valid to pad: Zero size")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:37
		// _ = "end of CoverTab[85491]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:38
		_go_fuzz_dep_.CoverTab[85492]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:38
		// _ = "end of CoverTab[85492]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:38
	// _ = "end of CoverTab[85486]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:38
	_go_fuzz_dep_.CoverTab[85487]++
													n := m - (len(b) % m)
													pb := make([]byte, len(b)+n)
													copy(pb, b)
													copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
													return pb, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:43
	// _ = "end of CoverTab[85487]"
}

// PKCS7Unpad removes RFC 2315 padding from byes where message size is m.
func PKCS7Unpad(b []byte, m int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:47
	_go_fuzz_dep_.CoverTab[85493]++
													if m <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:48
		_go_fuzz_dep_.CoverTab[85499]++
														return nil, errors.New("invalid message block size when unpadding")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:49
		// _ = "end of CoverTab[85499]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:50
		_go_fuzz_dep_.CoverTab[85500]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:50
		// _ = "end of CoverTab[85500]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:50
	// _ = "end of CoverTab[85493]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:50
	_go_fuzz_dep_.CoverTab[85494]++
													if b == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:51
		_go_fuzz_dep_.CoverTab[85501]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:51
		return len(b) == 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:51
		// _ = "end of CoverTab[85501]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:51
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:51
		_go_fuzz_dep_.CoverTab[85502]++
														return nil, errors.New("padded data not valid: Zero size")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:52
		// _ = "end of CoverTab[85502]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:53
		_go_fuzz_dep_.CoverTab[85503]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:53
		// _ = "end of CoverTab[85503]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:53
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:53
	// _ = "end of CoverTab[85494]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:53
	_go_fuzz_dep_.CoverTab[85495]++
													if len(b)%m != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:54
		_go_fuzz_dep_.CoverTab[85504]++
														return nil, errors.New("padded data not valid: Not multiple of message block size")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:55
		// _ = "end of CoverTab[85504]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:56
		_go_fuzz_dep_.CoverTab[85505]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:56
		// _ = "end of CoverTab[85505]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:56
	// _ = "end of CoverTab[85495]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:56
	_go_fuzz_dep_.CoverTab[85496]++
													c := b[len(b)-1]
													n := int(c)
													if n == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:59
		_go_fuzz_dep_.CoverTab[85506]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:59
		return n > len(b)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:59
		// _ = "end of CoverTab[85506]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:59
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:59
		_go_fuzz_dep_.CoverTab[85507]++
														return nil, errors.New("padded data not valid: Data may not have been padded")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:60
		// _ = "end of CoverTab[85507]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:61
		_go_fuzz_dep_.CoverTab[85508]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:61
		// _ = "end of CoverTab[85508]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:61
	// _ = "end of CoverTab[85496]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:61
	_go_fuzz_dep_.CoverTab[85497]++
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:62
		_go_fuzz_dep_.CoverTab[85509]++
														if b[len(b)-n+i] != c {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:63
			_go_fuzz_dep_.CoverTab[85510]++
															return nil, errors.New("padded data not valid")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:64
			// _ = "end of CoverTab[85510]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:65
			_go_fuzz_dep_.CoverTab[85511]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:65
			// _ = "end of CoverTab[85511]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:65
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:65
		// _ = "end of CoverTab[85509]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:66
	// _ = "end of CoverTab[85497]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:66
	_go_fuzz_dep_.CoverTab[85498]++
													return b[:len(b)-n], nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:67
	// _ = "end of CoverTab[85498]"
}

// GetHash generates the keyed hash value according to the etype's hash function.
func GetHash(pt, key []byte, usage []byte, etype etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:71
	_go_fuzz_dep_.CoverTab[85512]++
													k, err := etype.DeriveKey(key, usage)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:73
		_go_fuzz_dep_.CoverTab[85514]++
														return nil, fmt.Errorf("unable to derive key for checksum: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:74
		// _ = "end of CoverTab[85514]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:75
		_go_fuzz_dep_.CoverTab[85515]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:75
		// _ = "end of CoverTab[85515]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:75
	// _ = "end of CoverTab[85512]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:75
	_go_fuzz_dep_.CoverTab[85513]++
													mac := hmac.New(etype.GetHashFunc(), k)
													p := make([]byte, len(pt))
													copy(p, pt)
													mac.Write(p)
													return mac.Sum(nil)[:etype.GetHMACBitLength()/8], nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:80
	// _ = "end of CoverTab[85513]"
}

// GetChecksumHash returns a keyed checksum hash of the bytes provided.
func GetChecksumHash(b, key []byte, usage uint32, etype etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:84
	_go_fuzz_dep_.CoverTab[85516]++
													return GetHash(b, key, GetUsageKc(usage), etype)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:85
	// _ = "end of CoverTab[85516]"
}

// GetIntegrityHash returns a keyed integrity hash of the bytes provided.
func GetIntegrityHash(b, key []byte, usage uint32, etype etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:89
	_go_fuzz_dep_.CoverTab[85517]++
													return GetHash(b, key, GetUsageKi(usage), etype)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:90
	// _ = "end of CoverTab[85517]"
}

// VerifyChecksum compares the checksum of the msg bytes is the same as the checksum provided.
func VerifyChecksum(key, chksum, msg []byte, usage uint32, etype etype.EType) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:94
	_go_fuzz_dep_.CoverTab[85518]++

													expectedMAC, _ := GetChecksumHash(msg, key, usage, etype)
													return hmac.Equal(chksum, expectedMAC)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:97
	// _ = "end of CoverTab[85518]"
}

// GetUsageKc returns the checksum key usage value for the usage number un.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:100
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:100
// See RFC 3961 5.3 key-derivation function definition.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:103
func GetUsageKc(un uint32) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:103
	_go_fuzz_dep_.CoverTab[85519]++
													return getUsage(un, 0x99)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:104
	// _ = "end of CoverTab[85519]"
}

// GetUsageKe returns the encryption key usage value for the usage number un
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:107
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:107
// See RFC 3961 5.3 key-derivation function definition.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:110
func GetUsageKe(un uint32) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:110
	_go_fuzz_dep_.CoverTab[85520]++
													return getUsage(un, 0xAA)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:111
	// _ = "end of CoverTab[85520]"
}

// GetUsageKi returns the integrity key usage value for the usage number un
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:114
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:114
// See RFC 3961 5.3 key-derivation function definition.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:117
func GetUsageKi(un uint32) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:117
	_go_fuzz_dep_.CoverTab[85521]++
													return getUsage(un, 0x55)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:118
	// _ = "end of CoverTab[85521]"
}

func getUsage(un uint32, o byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:121
	_go_fuzz_dep_.CoverTab[85522]++
													var buf bytes.Buffer
													binary.Write(&buf, binary.BigEndian, un)
													return append(buf.Bytes(), o)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:124
	// _ = "end of CoverTab[85522]"
}

// IterationsToS2Kparams converts the number of iterations as an integer to a string representation.
func IterationsToS2Kparams(i uint32) string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:128
	_go_fuzz_dep_.CoverTab[85523]++
													b := make([]byte, 4, 4)
													binary.BigEndian.PutUint32(b, i)
													return hex.EncodeToString(b)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:131
	// _ = "end of CoverTab[85523]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:132
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/common/common.go:132
var _ = _go_fuzz_dep_.CoverTab
