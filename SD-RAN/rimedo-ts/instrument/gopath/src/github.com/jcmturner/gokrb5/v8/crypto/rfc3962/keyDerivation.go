//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:1
package rfc3962

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:1
)

import (
	"encoding/binary"
	"encoding/hex"
	"errors"

	"github.com/jcmturner/gofork/x/crypto/pbkdf2"
	"github.com/jcmturner/gokrb5/v8/crypto/etype"
)

const (
	s2kParamsZero = 4294967296
)

// StringToKey returns a key derived from the string provided according to the definition in RFC 3961.
func StringToKey(secret, salt, s2kparams string, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:17
	_go_fuzz_dep_.CoverTab[85745]++
														i, err := S2KparamsToItertions(s2kparams)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:19
		_go_fuzz_dep_.CoverTab[85747]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:20
		// _ = "end of CoverTab[85747]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:21
		_go_fuzz_dep_.CoverTab[85748]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:21
		// _ = "end of CoverTab[85748]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:21
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:21
	// _ = "end of CoverTab[85745]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:21
	_go_fuzz_dep_.CoverTab[85746]++
														return StringToKeyIter(secret, salt, i, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:22
	// _ = "end of CoverTab[85746]"
}

// StringToPBKDF2 generates an encryption key from a pass phrase and salt string using the PBKDF2 function from PKCS #5 v2.0
func StringToPBKDF2(secret, salt string, iterations int64, e etype.EType) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:26
	_go_fuzz_dep_.CoverTab[85749]++
														return pbkdf2.Key64([]byte(secret), []byte(salt), iterations, int64(e.GetKeyByteSize()), e.GetHashFunc())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:27
	// _ = "end of CoverTab[85749]"
}

// StringToKeyIter returns a key derived from the string provided according to the definition in RFC 3961.
func StringToKeyIter(secret, salt string, iterations int64, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:31
	_go_fuzz_dep_.CoverTab[85750]++
														tkey := e.RandomToKey(StringToPBKDF2(secret, salt, iterations, e))
														return e.DeriveKey(tkey, []byte("kerberos"))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:33
	// _ = "end of CoverTab[85750]"
}

// S2KparamsToItertions converts the string representation of iterations to an integer
func S2KparamsToItertions(s2kparams string) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:37
	_go_fuzz_dep_.CoverTab[85751]++
	//The s2kparams string should be hex string representing 4 bytes
	//The 4 bytes represent a number in big endian order
	//If the value is zero then the number of iterations should be 4,294,967,296 (2^32)
	var i uint32
	if len(s2kparams) != 8 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:42
		_go_fuzz_dep_.CoverTab[85754]++
															return int64(s2kParamsZero), errors.New("invalid s2kparams length")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:43
		// _ = "end of CoverTab[85754]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:44
		_go_fuzz_dep_.CoverTab[85755]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:44
		// _ = "end of CoverTab[85755]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:44
	// _ = "end of CoverTab[85751]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:44
	_go_fuzz_dep_.CoverTab[85752]++
														b, err := hex.DecodeString(s2kparams)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:46
		_go_fuzz_dep_.CoverTab[85756]++
															return int64(s2kParamsZero), errors.New("invalid s2kparams, cannot decode string to bytes")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:47
		// _ = "end of CoverTab[85756]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:48
		_go_fuzz_dep_.CoverTab[85757]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:48
		// _ = "end of CoverTab[85757]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:48
	// _ = "end of CoverTab[85752]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:48
	_go_fuzz_dep_.CoverTab[85753]++
														i = binary.BigEndian.Uint32(b)
														return int64(i), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:50
	// _ = "end of CoverTab[85753]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:51
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/keyDerivation.go:51
var _ = _go_fuzz_dep_.CoverTab
